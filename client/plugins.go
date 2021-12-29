package main

import (
	_ "github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	_ "github.com/asim/go-micro/plugins/registry/etcd/v4"
	_ "github.com/asim/go-micro/plugins/registry/kubernetes/v4"
	_ "github.com/asim/go-micro/plugins/transport/nats/v4"
)
