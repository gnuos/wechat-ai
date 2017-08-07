BinPath = $(shell `pwd`)
Service = wechat.service

.PHONY: all
all: wechat-ai

.PHONY: wechat-ai
wechat-ai:
	go build

.PHONY: service
service:
	sed -i "s/{{ wechat-ai path }}/$(BinPath)/" wechat.service
	cp -f $(Service) /lib/systemd/system/

.PHONY: clean
clean:
	go clean
	rm -f log/*.log

