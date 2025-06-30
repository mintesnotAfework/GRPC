package main

import (
	"net"

	"github.com/mintesnotAfework/GRPC/GRPC/number"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	number.RegisterNumServiceServer(s, &number.NumServer{})
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
