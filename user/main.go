package main

import (
	"user/config"
	"user/core"
	"user/service"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {

	config.Init()
	// 注册etcd
	etcdRegi := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	// 得到微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),    // 名字
		micro.Address("127.0.0.1:8082"), // 地址
		micro.Registry(etcdRegi),        // 注册到etcd 中
	)

	microService.Init()
	// 服务注册
	_ = service.RegisterUserServiceHandler(microService.Server(), new(core.UserService))

	// 启动微服务
	microService.Run()
}
