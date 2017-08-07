package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/levigross/grequests"
)

// 机器人回复的数据的json绑定对象
type aiReply struct {
	code int
	Text string `json:"text"`
}

func tlAI(info string) string {
	tlURL := config.Ai.ApiUrl
	ro := &grequests.RequestOptions{
		Params: map[string]string{
			"key":  config.Ai.ApiKey,
			"info": info,
		},
	}

	r, err := grequests.Get(tlURL, ro)
	if err != nil {
		log.Println(err)
		return ""
	}

	defer r.Close()
	reply := new(aiReply)
	json.Unmarshal([]byte(r.String()), reply)
	return reply.Text
}

// 对机器人API发送请求
func ask(info string) string {
	URL := fmt.Sprintf(config.Ai.ApiUrl+"?key=%s&info=%s",
		config.Ai.ApiKey, url.QueryEscape(info))
	resp, err := http.Get(URL)
	if err != nil {
		log.Println(err)
		return config.Ai.Greeting
	}
	defer resp.Body.Close()
	reply := new(aiReply)
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(reply)
	return reply.Text
}
