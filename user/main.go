//go:build !js && !wasm
// +build !js,!wasm

package main

import (
	"user/config"
	"user/core"
	"user/services"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {
	config.Init()

	// etcd 注册件
	etcReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	// 获得微服务实例
	microService := micro.NewService(
		micro.Name("rpcUserService"),    // 名字
		micro.Address("127.0.0.1:8082"), // 地址
		micro.Registry(etcReg),          // 注册件
	)

	// 接收命令行参数
	microService.Init()

	// 服务注册
	_ = services.RegisterUserServiceHandler(microService.Server(), new(core.UserService))

	// 启动微服务
	_ = microService.Run()

}
