package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mintesnotAfework/GRPC/GRPC/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("connecting.....")
	conn, err := grpc.NewClient("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer conn.Close()
	fmt.Println("adding new service.....")
	c := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	fmt.Println("Fetching user.....")
	r, err := c.GetUser(ctx, &user.UserRequest{UserId: "12"})
	if err != nil {
		panic(err)
	}

	fmt.Println("Client received:", r.String())
}
