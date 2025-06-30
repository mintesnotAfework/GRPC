package main

import (
	"fmt"

	"github.com/mintesnotAfework/protocol-buffer/group"
	"google.golang.org/protobuf/proto"
)

func main() {
	u1 := group.Group_User{UserId: "1", Email: "hello@world.com"}
	u2 := group.Group_User{UserId: "2", Email: "non@world.com"}

	g := group.Group{
		Score:    3.2,
		User:     []*group.Group_User{&u1, &u2},
		Category: group.Category_OPERATOR,
		Id:       1,
	}
	fmt.Println(g.String())

	encoded, err := proto.Marshal(&g)
	if err != nil {
		panic("can not marshal")
	}

	v := group.Group{}
	err = proto.Unmarshal(encoded, &v)
	if err != nil {
		panic("can not unmarshal")
	}

	fmt.Println("Recovered:", v.String())
}
