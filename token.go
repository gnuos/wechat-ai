package main

import (
	"encoding/json"

	"github.com/levigross/grequests"
)

type Token struct {
	AccessToken string `json:"access_token"`
	Expire      string `json:"expires_in"`
}

// 获取access_token
func (t *Token) Get() []byte {
	var args = map[string]string{
		"appid":      config.Wx.AppID,
		"secret":     config.Wx.AppSecret,
		"grant_type": "client_credential",
	}

	ro := &grequests.RequestOptions{
		Params: args,
	}

	res, _ := grequests.Get("https://api.weixin.qq.com/cgi-bin/token", ro)
	return res.Bytes()
}

func (t *Token) Parse(at []byte) (err error) {
	if err = json.Unmarshal(at, t); err != nil {
		return err
	}

	return
}
