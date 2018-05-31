package main

import (
	"log"
	"net/http"

	"github.com/devfeel/dotweb"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

// 路由 HEAD("*", Null)
func Null(ctx dotweb.Context) error {
	return nil
}

// 路由 GET("/ip", GetIP)
func GetIP(ctx dotweb.Context) error {
	var resp struct {
		IP string `json:"ip"`
	}
	resp.IP = ctx.Request().QueryHeader("X-Forwarded-For")

	return ctx.WriteJson(&resp)
}

// 路由 Any("/", Default)
func Default(ctx dotweb.Context) error {
	// 从配置文件获取微信的AppID和其他参数
	var wxConf = &wechat.Config{
		AppID:          config.Wx.AppID,
		AppSecret:      config.Wx.AppSecret,
		Token:          config.Wx.Token,
		EncodingAESKey: config.Wx.EncodingAESKey,
	}

	// 检查请求来源
	if !validateUrl(ctx) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return ctx.WriteStringC(http.StatusForbidden, `{"message": 403}`)
	}

	if echostr := ctx.QueryString("echostr"); echostr != "" {
		return ctx.WriteString(echostr)
	}

	wc := wechat.NewWechat(wxConf)
	server := wc.GetServer(ctx.Request().Request, ctx.Response().Writer())

	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		text := message.NewText(tlAI(msg.Content))
		return &message.Reply{message.MsgTypeText, text}
	})

	//处理消息接收以及回复
	err := server.Serve()
	if err != nil {
		log.Println(err)
		return err
	}
	//发送回复的消息
	return server.Send()
}

func InitRoute(server *dotweb.HttpServer) {
	server.Any("/", Default)
	server.GET("/ip", GetIP)
	server.HEAD("/ip", Null)
	server.GET("/favicon.ico", func(ctx dotweb.Context) error {
		return ctx.File("favicon.ico")
	})
}
