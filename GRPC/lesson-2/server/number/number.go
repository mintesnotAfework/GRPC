package number

import (
	"errors"
	"fmt"
	"math/rand"
)

type NumServer struct {
	UnimplementedNumServiceServer
}

func (n *NumServer) Rnd(req *NumRequest, stream NumService_RndServer) error {
	fmt.Println(req.String())
	if req.N <= 0 {
		return errors.New("n must be greater than zero")
	}
	if req.To <= req.From {
		return errors.New("to must be greater or equal than from")
	}

	done := make(chan bool)
	go func() {
		for counter := 0; counter < int(req.N); counter++ {
			i := rand.Intn(int(req.To)-int(req.From)+1) + int(req.To)
			resp := NumResponse{I: int64(i), Remaining: req.N - int64(counter)}
			stream.Send(&resp)
		}
		done <- true
	}()

	<-done
	return nil
}
