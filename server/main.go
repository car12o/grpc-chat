package main

import (
	"fmt"
	"log"
	"net"

	"github.com/car12o/grpc-chat/proto"
	"github.com/car12o/grpc-chat/server/chat"
	"google.golang.org/grpc"
)

const port = 3000

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterChatServer(grpcServer, chat.NewChatServer())
	log.Printf("Serven listen on port: %d\n", port)
	grpcServer.Serve(lis)
}
