package main

import (
	"net"

	"github.com/mintesnotAfework/GRPC/GRPC/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	chat.RegisterChatServiceServer(s, &chat.ChatServer{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
