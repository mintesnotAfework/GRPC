package main

import (
	"fmt"

	"github.com/mintesnotAfework/protocol-buffer/user"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func main() {
	u := user.User{
		UserId: "1",
		Email:  "hello@email.com",
		Info:   []*anypb.Any{{Value: []byte("john doe"), TypeUrl: "urltype"}},
	}

	fmt.Println(u.String())

	encoded, err := proto.Marshal(&u)
	if err != nil {
		panic("can not marshal")
	}

	v := user.User{}
	err = proto.Unmarshal(encoded, &v)
	if err != nil {
		panic("can not unmarshal")
	}

	fmt.Println("Recovered:", v.String())
}
