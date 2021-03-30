FROM golang:1.16-stretch
RUN apt-get update && apt-get install -y unzip
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.15.6/protoc-3.15.6-linux-x86_64.zip -P /usr/local/protoc && \
  unzip /usr/local/protoc/protoc-3.15.6-linux-x86_64.zip -d /usr/local/protoc && \
  ln -s /usr/local/protoc/bin/protoc -t /usr/local/bin
RUN go get -u google.golang.org/protobuf/cmd/protoc-gen-go \
  google.golang.org/grpc/cmd/protoc-gen-go-grpc
RUN wget https://github.com/grpc/grpc-web/releases/download/1.2.1/protoc-gen-grpc-web-1.2.1-linux-x86_64 -P /usr/local/protoc && \
  mv /usr/local/protoc/protoc-gen-grpc-web-1.2.1-linux-x86_64 /usr/local/protoc/bin/protoc-gen-grpc-web && \
  ln -s /usr/local/protoc/bin/protoc-gen-grpc-web -t /usr/local/bin && \
  chmod u=rwx,g=rx,o= /usr/local/bin/protoc-gen-grpc-web
COPY proto /grpc-chat/proto
COPY Makefile /grpc-chat
WORKDIR /grpc-chat
RUN make proto
