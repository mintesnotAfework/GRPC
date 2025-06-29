package main

import (
	"fmt"

	"github.com/mintesnotAfework/protocol-buffer/user"
	"google.golang.org/protobuf/proto"
)

func main() {
	u := user.User{
		UserId: "1",
		Email:  "hello@world.com",
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
