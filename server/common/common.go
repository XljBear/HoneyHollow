package common

import (
	"HoneyHollow/server/globalStructs"
	"math/rand"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

func RandString(length int, symbols bool) string {
	if length < 1 {
		return ""
	}
	var char string
	if symbols {
		char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	} else {
		char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	}
	charArr := strings.Split(char, "")
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))

	l := len(charArr)
	for i := l - 1; i > 0; i-- {
		r := ran.Intn(i)
		charArr[r], charArr[i] = charArr[i], charArr[r]
	}
	randChar := charArr[:length]
	return strings.Join(randChar, "")
}

func CountStringLen(str string) int {
	return utf8.RuneCountInString(str)
}

func MakeResponse(statusCode int, message string, data interface{}) *globalStructs.CommonResponseStruct {
	resp := &globalStructs.CommonResponseStruct{}
	resp.Code = statusCode
	resp.Message = message
	resp.Data = data
	return resp
}

func IsValidURL(url string) bool {
	// 正则表达式用于匹配网址
	pattern := `^(https?|ftp)://[^\s/$.?#].[^\s]*$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(url)
}
