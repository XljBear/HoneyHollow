package job

import (
	"HoneyHollow/server/services"
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2/log"
	"io"
	"net/http"
	"net/url"
	"path"
	"time"
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
			log.Info(bookmarks.Url)
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
