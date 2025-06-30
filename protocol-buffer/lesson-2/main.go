package main

import (
	"fmt"

	"github.com/mintesnotAfework/protocol-buffer/user"
	"google.golang.org/protobuf/proto"
)

func main() {
	u1 := user.User{
		UserId: "1",
		Email:  "hello@world.com",
	}
	u2 := user.User{
		UserId: "2",
		Email:  "hola@world.com",
	}

	g := user.Group{
		Id:       1,
		Score:    3.2,
		Users:    []*user.User{&u1, &u2},
		Category: user.Category_DEVELOPER,
	}

	fmt.Println(g.String())

	encoded, err := proto.Marshal(&g)
	if err != nil {
		panic("can not marshal")
	}

	v := user.Group{}
	err = proto.Unmarshal(encoded, &v)
	if err != nil {
		panic("can not unmarshal")
	}

	fmt.Println("Recovered:", v.String())
}
