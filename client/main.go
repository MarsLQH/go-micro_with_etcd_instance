package main

import (
	pb "WithEtcd/proto"
	"context"
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
)

type Login struct {
	Client pb.UserService
}


const (
	ETCD_ADDRESS  = "8.142.86.204:2379"
)

func main() {

	registry := etcd.NewRegistry(registry.Addrs(ETCD_ADDRESS))
	service := micro.NewService(
		micro.Registry(registry),
	)
	service.Init()
	c := service.Client()
	// create request/response
	req := pb.LoginRequest{Name: "Mars"}
	request := client.NewRequest("micro.srv.user", "User.Login", &req)
	response := new(pb.LoginResponse)
	// call service
	err := c.Call(context.Background(), request, response)

	log.Error(err)
	log.Info(response)
}
