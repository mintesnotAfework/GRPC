package main

import (
	"fmt"
	"net"

	"github.com/mintesnotAfework/GRPC/GRPC/number"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("starting server")
	listener, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		panic(err)
	}

	fmt.Println("registering service....")
	s := grpc.NewServer()
	number.RegisterNumServiceServer(s, &number.NumServer{}) // registering server

	fmt.Println("listening to incoming request....")
	if err := s.Serve(listener); err != nil { // listening to incoming request
		panic(err)
	}
}
