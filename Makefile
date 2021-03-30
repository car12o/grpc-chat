.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/chat.proto
	protoc --js_out=import_style=commonjs:. \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:. \
		proto/chat.proto

dev:
	GO111MODULE=off go get -u github.com/cosmtrek/air
	air -c server/air.toml

test:
	go test -v ./server/...
