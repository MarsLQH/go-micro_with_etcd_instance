package handler

import (
	pb "WithEtcd/proto"
	"context"
	"fmt"
)

type User struct {
}

func (u User) Login(ctx context.Context, request *pb.LoginRequest, response *pb.LoginResponse) error {
	fmt.Println("receive request")
	response.Uid = 1
	response.Token = "123456789"
	return nil
}
