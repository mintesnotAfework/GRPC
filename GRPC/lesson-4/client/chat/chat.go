package chat

import (
	"fmt"
	"time"
)

func Chat(stream ChatService_SendTxtClient, done chan bool) {
	t := time.NewTicker(time.Millisecond * 500)
	for {
		select {
		case <-done:
			return
		case <-t.C:
			err := stream.Send(&ChatRequest{Txt: "Hello", Id: 1, To: 2})
			if err != nil {
				panic(err)
			}
		}
	}
}

func Stats(stream ChatService_SendTxtClient, done chan bool) {
	for {
		stats, err := stream.Recv()
		if err != nil {
			panic(err)
		}

		fmt.Println(stats.String())
		if stats.TotalChar > 35 {
			fmt.Println("beyond the limit!!!")
			done <- true
			stream.CloseSend()
			return
		}
	}
}
