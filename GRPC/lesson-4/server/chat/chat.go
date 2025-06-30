package chat

import (
	"fmt"
	"io"
	"time"
)

type ChatServer struct {
	UnimplementedChatServiceServer
}

func (c *ChatServer) SendTxt(stream ChatService_SendTxtServer) error {
	var total int64 = 0
	go func() {
		for {
			t := time.NewTicker(time.Second * 2)
			<-t.C
			stream.Send(&StatsResponse{TotalChar: total})
		}

	}()
	for {
		next, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("client closed")
			return nil
		}

		if err != nil {
			return err
		}
		total = total + int64(len(next.Txt))
		fmt.Println("->", next.Txt)
	}

}
