FROM golang:1.16-stretch
RUN apt-get update && apt-get install -y unzip xz-utils
# Install potroc
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.15.6/protoc-3.15.6-linux-x86_64.zip -P /usr/local/protoc && \
  unzip /usr/local/protoc/protoc-3.15.6-linux-x86_64.zip -d /usr/local/protoc && \
  ln -s /usr/local/protoc/bin/protoc -t /usr/local/bin
# Install grpc Go plugins
RUN go get -u google.golang.org/protobuf/cmd/protoc-gen-go \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc
# Install grpc web plugin
RUN wget https://github.com/grpc/grpc-web/releases/download/1.2.1/protoc-gen-grpc-web-1.2.1-linux-x86_64 -P /usr/local/protoc && \
  mv /usr/local/protoc/protoc-gen-grpc-web-1.2.1-linux-x86_64 /usr/local/protoc/bin/protoc-gen-grpc-web && \
  ln -s /usr/local/protoc/bin/protoc-gen-grpc-web -t /usr/local/bin && \
  chmod u=rwx,g=rx,o= /usr/local/bin/protoc-gen-grpc-web
# Install Node
RUN wget https://nodejs.org/dist/v14.16.0/node-v14.16.0-linux-x64.tar.xz -P /usr/local/node && \
  tar -xf /usr/local/node/node-v14.16.0-linux-x64.tar.xz -C /usr/local/node --strip-components=1
ENV PATH="${PATH}:/usr/local/node/bin"
# Install grpc Node plugins
RUN npm i -g --unsafe-perm grpc grpc-tools grpc_tools_node_protoc_ts
COPY proto /grpc-chat/proto
COPY Makefile /grpc-chat
WORKDIR /grpc-chat
CMD [ "make", "proto" ]
