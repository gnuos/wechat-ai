Bin = wechat-ai
BinDir = $(shell pwd)
Service = wechat.service

.PHONY: all
all: wechat-ai

.PHONY: wechat-ai
wechat-ai:
	go fmt && go build

.PHONY: service
service:
	sed -i "s|{{ wechat-ai path }}|$(BinDir)|" $(Service)
	@cp -f $(Service) /lib/systemd/system/

.PHONY: clean
clean:
	go clean
	@rm -f log/*.log

