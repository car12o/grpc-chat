# proto
.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/chat.proto
	protoc --js_out=import_style=commonjs:. \
		--grpc-web_out=import_style=commonjs,mode=grpcwebtext:. \
		proto/chat.proto
	grpc_tools_node_protoc \
    --plugin=protoc-gen-ts=/usr/local/node/bin/protoc-gen-ts \
    --ts_out=./proto \
    -I ./proto \
    proto/*.proto

# server
sv.dev:
	GO111MODULE=off go get -u github.com/cosmtrek/air
	air -c server/air.toml

sv.test:
	go test -v ./server/...

# client
cl.proto:
	cp ./proto/*.js ./client/src/proto
	cp ./proto/*.ts ./client/src/proto

cl.serve:
	npm run start --prefix ./client

cl.dev: cl.proto cl.serve
