gen-buf:
	buf generate proto
.PHONY: gen

gen-protoc:
	protoc -I ./proto \
	--go_out ./gen/go --go_opt paths=source_relative \
	--go-grpc_out ./gen/go --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./gen/go --grpc-gateway_opt paths=source_relative \
	./proto/exp.proto


proxy:
	go run proxy/proxy.go
.PHONY: proxy

server:
	go run server/server.go
.PHONY: server
