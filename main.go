package main

import (
	"flag"
	"fmt"

	"github.com/devfeel/dotweb"
	"github.com/devfeel/middleware/accesslog"
)

var config *Config

func main() {
	var configFile string
	var debug bool
	var err error

	// 从config.json文件中获取配置参数
	flag.StringVar(&configFile, "c", "config.json", "specify config file")
	flag.BoolVar(&debug, "d", false, "debug mode")
	flag.Parse()
	config, err = ParseConfig(configFile)
	if err != nil {
		panic("a vailid json config file must exist")
	}

	app := dotweb.New()

	app.Use(accesslog.Middleware())

	app.SetEnabledLog(true)

	if config.LogPath != "" {
		app.SetLogPath(config.LogPath)

		if debug {
			app.SetPProfConfig(true, 8081)
		}
	} else {
		app.SetDevelopmentMode()
	}

	// 设置路由
	InitRoute(app.HttpServer)

	// 启动Dotweb
	err = app.StartServer(config.Port)
	fmt.Println("dotweb.StartServer error => ", err)
}
