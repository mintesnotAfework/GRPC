package main

import (
	"context"

	"github.com/mintesnotAfework/GRPC/GRPC/chat"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := chat.NewChatServiceClient(conn)
	stream, err := c.SendTxt(context.Background())
	if err != nil {
		panic(err)
	}

	done := make(chan bool)
	go chat.Stats(stream, done)
	go chat.Chat(stream, done)
	<-done
}
