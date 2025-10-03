package main

import (
	"api-gateway/service"
	"api-gateway/weblib"
	"api-gateway/wrappers"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-micro/v2/web"
)

func main() {

	etcdRegi := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	// 注册用户服务
	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)

	// 用户服务调用实例
	userService := service.NewUserService("rpcUserService", userMicroService.Client())

	// task
	taskMicroService := micro.NewService(
		micro.Name("taskService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)

	taskService := service.NewTaskService("rpcTaskService", taskMicroService.Client())

	server := web.NewService(
		web.Name("httpService"),
		web.Address("127.0.0.1:4000"),
		web.Handler(weblib.NewRouter(userService, taskService)),
		web.Registry(etcdRegi),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocel": "http"}),
	)

	_ = server.Init()
	_ = server.Run()
}
