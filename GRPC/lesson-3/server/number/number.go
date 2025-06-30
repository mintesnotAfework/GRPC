package number

import "io"

type NumServer struct {
	UnimplementedNumServiceServer
}

func (n *NumServer) Sum(stream NumService_SumServer) error {
	var total int64 = 0
	var counter int = 0
	for {
		next, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&NumResponse{I: total})
			return nil
		}
		if err != nil {
			panic(err)
		}
		total = total + next.X
		counter++
	}
}
