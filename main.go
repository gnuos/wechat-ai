package main

import (
	"flag"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var config *Config

// 定义Web服务的Server头
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "WeChat")
		return next(c)
	}
}

func main() {
	var configFile string
	var debug uint
	var err error

	// 从config.json文件中获取配置参数
	flag.StringVar(&configFile, "c", "config.json", "specify config file")
	flag.UintVar(&debug, "d", 0, "debug mode")
	flag.Parse()
	config, err = ParseConfig(configFile)
	if err != nil {
		panic("a vailid json config file must exist")
	}

	web := echo.New()

	web.Use(ServerHeader)

	// 设置日志文件格式
	loggerConfig := middleware.LoggerConfig{
		Format: config.LogFormat,
	}
	if config.LogFile != "" {
		f, err := os.OpenFile(config.LogFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		loggerConfig.Output = f
		web.Logger.SetOutput(f)

		if debug == 1 {
			web.Debug = true
		}
	} else {
		web.Debug = true
	}

	web.Use(middleware.LoggerWithConfig(loggerConfig))
	web.Use(middleware.Recover())

	if web.Debug {
		web.Logger.SetLevel(log.DEBUG)
	}

	// 定义跨域请求所需授权的域和方法
	web.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.POST, echo.DELETE},
	}))

	// 定义路由
	web.HEAD("/*", Null)
	web.Any("/", Default)
	web.GET("/ip", GetIP)
	web.GET("/favicon.ico", func(ctx echo.Context) error {
		return ctx.File("favicon.ico")
	})

	// 启动Web服务
	web.Start(config.Listen)
}
