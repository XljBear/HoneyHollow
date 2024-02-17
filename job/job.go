package job

import (
	"HoneyHollow/server/services"
	"bytes"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
	"unicode/utf8"
)

func getHostURL(websiteURL string) string {
	baseURL, err := url.Parse(websiteURL)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%s://%s", baseURL.Scheme, baseURL.Host)
}
func getAbsoluteFaviconURL(websiteURL, faviconURL string) string {
	// 解析网址
	baseURL, err := url.Parse(websiteURL)
	if err != nil {
		return ""
	}

	// 如果 faviconURL 是绝对路径，则直接返回
	faviconParsedURL, err := url.Parse(faviconURL)
	if err != nil || faviconParsedURL.IsAbs() {
		return faviconURL
	}

	// 如果 faviconURL 是相对路径，则拼接网站的主域名和路径
	return fmt.Sprintf("%s://%s", baseURL.Scheme, path.Join(baseURL.Host, faviconURL))
}
func checkFaviconAccessibility(faviconURL string) bool {
	// 发送 HTTP HEAD 请求
	resp, err := http.Head(faviconURL)
	if err != nil {
		// 请求失败，返回 false
		return false
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	// 检查响应状态码
	if resp.StatusCode == http.StatusOK {
		return true
	}

	return false
}
func catchWebInfo(url string) (title string, faviconURL string) {
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3.1 Safari/605.1.15"
	c.OnHTML("link[rel='icon']", func(e *colly.HTMLElement) {
		faviconURL = e.Attr("href")
	})
	c.OnHTML("title", func(e *colly.HTMLElement) {
		title = e.Text
		if !utf8.ValidString(title) && isGBKEncoding([]byte(title)) {
			utf8Title, err := gbkToUTF8([]byte(title))
			if err == nil {
				title = string(utf8Title)
			}
		}
	})
	_ = c.Visit(url)
	if faviconURL == "" {
		normalFaviconURL := getHostURL(url) + "/favicon.ico"
		if checkFaviconAccessibility(normalFaviconURL) {
			faviconURL = normalFaviconURL
		}
	} else {
		faviconURL = getAbsoluteFaviconURL(url, faviconURL)
	}
	return
}
func ProcessBookmarks() {
	for {
		needProcessBookmarks := services.GetAllNeedProcessBookmarks()
		for _, bookmarks := range needProcessBookmarks {
			title, faviconURL := catchWebInfo(bookmarks.Url)
			bookmarks.Title = title
			bookmarks.Icon = faviconURL
			bookmarks.NeedProcess = false
			err := services.UpdateBookmarks(&bookmarks)
			if err != nil {
				log.Info(err)
			}
			time.Sleep(time.Second)
		}
		time.Sleep(time.Second)
	}
}

// 判断文本是否是GBK编码
func isGBKEncoding(data []byte) bool {
	for i := 0; i < len(data)-1; i++ {
		if data[i] >= 0x81 && data[i] <= 0xFE {
			if data[i+1] >= 0x40 && data[i+1] <= 0xFE {
				return true
			}
		}
	}
	return false
}

// 将GBK编码的文本转换为UTF-8
func gbkToUTF8(gbkContent []byte) ([]byte, error) {
	reader := simplifiedchinese.GBK.NewDecoder().Reader(bytes.NewReader(gbkContent))
	utf8Content, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return utf8Content, nil
}
