package main

import (
	"github.com/levigross/grequests"
)

// 获取access_token
func GetAccessToken() string {
	var args = map[string]string{
		"appid":      config.Wx.AppID,
		"secret":     config.Wx.AppSecret,
		"grant_type": "client_credential",
	}

	ro := &grequests.RequestOptions{
		Params: args,
	}

	res, _ := grequests.Get("https://api.weixin.qq.com/cgi-bin/token", ro)
	return res.String()
}
