package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// 配置文件的json绑定对象
type Config struct {
	Listen    string `json:"listen"`
	LogFile   string `json:"log_file"`
	LogFormat string `json:"log_format"`
	Wx        Wexin  `json:"weixin"`
	Ai        AI     `json:"ai"`
}

type Wexin struct {
	AppID          string `json:"appid"`
	AppSecret      string `json:"secret"`
	Token          string `json:"token"`
	EncodingAESKey string `json:"encodingAESKey"`
}

type AI struct {
	Greeting string `json:"greeting"`
	ApiUrl   string `json:"url"`
	ApiKey   string `json:"key"`
}

// 获取配置文件中的配置参数
func ParseConfig(path string) (config *Config, err error) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	config = &Config{}
	if err = json.Unmarshal(data, config); err != nil {
		return nil, err
	}

	return
}
