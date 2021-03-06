package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"
	"strings"

	"github.com/labstack/echo"
)

// 对请求的参数做SHA1校验
func makeSignature(timestamp, nonce string) string {
	sl := []string{config.Wx.Token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

// 检查url中的参数
func validateUrl(ctx echo.Context) bool {
	timestamp := ctx.QueryParam("timestamp")
	nonce := ctx.QueryParam("nonce")
	signatureGen := makeSignature(timestamp, nonce)
	signatureIn := ctx.QueryParam("signature")

	if signatureGen != signatureIn {
		return false
	}

	return true
}
