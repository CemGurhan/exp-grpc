gen-buf:
	buf generate proto
.PHONY: gen-buf

gen-protoc:
	protoc -I ./proto \
	--go_out ./gen/go --go_opt paths=source_relative \
	--go-grpc_out ./gen/go --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./gen/go --grpc-gateway_opt paths=source_relative \
	./proto/exp.proto

gen-protoc-openapi:
	protoc -I ./proto \
	--go_out ./gen/go --go_opt paths=source_relative \
	--go-grpc_out ./gen/go --go-grpc_opt paths=source_relative \
	--grpc-gateway_out ./gen/go --grpc-gateway_opt paths=source_relative \
	--openapiv2_out . \
    --openapiv2_opt logtostderr=true \
	--openapiv2_opt use_go_templates=true \
	./proto/exp.proto


proxy:
	go run proxy/proxy.go
.PHONY: proxy

server:
	go run server/server.go
.PHONY: server
