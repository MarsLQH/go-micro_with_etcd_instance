package main

import (
	pb "WithEtcd/proto"
	"WithEtcd/service/handler"
	"go-micro.dev/v4"
	microRegistry "go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"log"
	"time"

	//"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	//"go-micro.dev/v4/cmd/protoc-gen-micro/plugin/micro"
)

const (
	ETCD_ADDRESS  = "x.x.x.x:2379"
	MICRO_ADDRESS = "x.x.x.x:5003"
)

func main() {
	registry := etcd.NewRegistry(microRegistry.Addrs(ETCD_ADDRESS))

	service := micro.NewService(
		micro.Name("micro.srv.user"),
		micro.Registry(registry),
		micro.Address(MICRO_ADDRESS),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)
	service.Init()
	service.Server().Init(server.Wait(nil))
	pb.RegisterUserHandler(
		service.Server(),
		new(handler.User),
	)
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
