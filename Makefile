gen:
	buf generate proto
.PHONY: gen

proxy:
	go run proxy/proxy.go
.PHONY: proxy

server:
	go run server/server.go
.PHONY: server
