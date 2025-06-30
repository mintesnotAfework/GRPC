package main

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/mintesnotAfework/GRPC/GRPC/lesson-2/client/number"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	c := number.NewNumServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()

	stream, err := c.Rnd(ctx, &number.NumRequest{N: 15, From: 0, To: 200})
	if err != nil {
		panic(err)
	}

	done := make(chan bool)
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				done <- true
				return
			}
			if err != nil {
				panic(err)
			}
			fmt.Println("Received:", resp.String())
		}
	}()
	<-done
	fmt.Println("Client done")
}
