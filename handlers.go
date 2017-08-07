package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/silenceper/wechat"
	"github.com/silenceper/wechat/message"
)

// 路由 web.HEAD("/*", Null)
func Null(ctx echo.Context) error {
	return nil
}

// 路由 web.GET("/ip", GetIP)
func GetIP(ctx echo.Context) error {
	var resp struct {
		IP string `json:"ip"`
	}
	resp.IP = ctx.RealIP()

	return ctx.JSON(http.StatusOK, &resp)

}

// 路由 web.Any("/", GetToken)
// 在验证微信公众号服务器时，需要将默认的路由改为这个
func GetToken(ctx echo.Context) error {
	if !validateUrl(ctx) {
		log.Println("Wechat Service: this http request is not from Wechat platform!")
		return ctx.String(http.StatusMethodNotAllowed, `{"message": 405}`)
	}

	echostr := ctx.QueryParam("echostr")
	return ctx.String(http.StatusOK, echostr)
}

// 路由 web.Any("/", Default)
func Default(ctx echo.Context) error {
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
		return ctx.String(http.StatusForbidden, `{"message": 403}`)
	}

	wc := wechat.NewWechat(wxConf)
	server := wc.GetServer(ctx.Request(), ctx.Response().Writer)

	//设置接收消息的处理方法
	server.SetMessageHandler(func(msg message.MixMessage) *message.Reply {
		text := message.NewText(ask(msg.Content))
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
