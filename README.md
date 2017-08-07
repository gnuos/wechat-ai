# WeChat-AI

微信公众号聊天机器人，类似于微信小冰的功能

## 项目特点

- 使用了Echo框架作为Web服务
- 使用了[silenceper/wechat](https://github.com/silenceper/wechat)微信SDK做微信消息处理
- 使用了图灵机器人网站的AI接口
- 使用JSON格式保存配置文件


## 安装

```bash
git clone https://github.com/gnuos/wechat-ai
cd wechat-ai
go build
```

## 快速开始

在安装之后，会生成一个wechat-ai文件，修改config.json配置文件，把相关的参数的值替换为你自己的环境，然后在项目目录中运行./wechat-ai，就启动成功了。
如果你用的操作系统是通过systemd管理服务的，可以将 wechat.service 文件复制到/lib/systemd/system/目录下面，然后将文件中的wechat-ai执行路径修改为你的路径，再执行systemctl daemon-reload命令重新加载服务文件，就可以通过systemd管理wechat-ai了。


## 开发文档

由于本项目主要是基于[Echo](https://echo.labstack.com/)和[silenceper/wechat](https://github.com/silenceper/wechat)框架的，如果要修改本项目的代码或者做研究学习，请参考以上两个项目的文档。


## License

Apache License, Version 2.0


