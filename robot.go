package main

import (
	"encoding/json"
	"log"

	"github.com/levigross/grequests"
)

// 机器人回复的数据的json绑定对象
type aiReply struct {
	code int
	Text string `json:"text"`
}

// 对机器人API发送请求
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
