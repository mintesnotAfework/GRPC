package main

import (
	"context"
	"fmt"
	"time"

	"github.com/mintesnotAfework/GRPC/GRPC/number"
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

	stream, err := c.Sum(ctx)
	if err != nil {
		panic(err)
	}

	from, to := 1, 200
	for i := from; i <= to; i++ {
		err = stream.Send(&number.NumRequest{X: int64(i)})
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("waiting for response")
	result, err := stream.CloseAndRecv()
	if err != nil {
		panic(err)
	}
	fmt.Println(result.String())
}
