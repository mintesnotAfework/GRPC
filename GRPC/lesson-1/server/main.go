package main

import (
	"fmt"
	"net"

	"github.com/mintesnotAfework/GRPC/GRPC/user"
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
	user.RegisterUserServiceServer(s, &user.UserServer{}) // registering server

	fmt.Println("listening to incoming request....")
	if err := s.Serve(listener); err != nil { // listening to incoming request
		panic(err)
	}
}
