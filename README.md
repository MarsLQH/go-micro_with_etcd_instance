# go-micro_with_etcd_instance
An instance of go-micro with etcd。最新版go-micro用Etcd做服务发现和注册的实例

对比几大Golang微服务框架,go-micro具有开发简洁,易于维护,功能齐全的特性。  
做分布式微服务少不了服务注册与发现，本文基于最新版本的go-micro，实现用etcd做服务注册与发现。  

注：
etcd是以插件的形式用来做registry的  
使用方法如下：  
服务端:  
go build -o service ./service/main.go  ./service/plugins.go  
./service --registry=etcd  --registry_address=x.x.x.x:2379 --server_advertise x.x.x.x:5003

客户端：  
go build -o client ./client/main.go  ./service/plugins.go  

生成protobuf：  
protoc --micro_out=./ --go_out=./  ./proto/user.proto  


