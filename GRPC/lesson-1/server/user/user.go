package user

import (
	"context"
	"fmt"
)

type UserServer struct {
	UnimplementedUserServiceServer
}

func (u *UserServer) GetUser(ctx context.Context, req *UserRequest) (*User, error) {
	fmt.Println("server recieved: ", req.String())
	return &User{UserId: "1", Email: "hello@world.com"}, nil
}
