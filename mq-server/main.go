package main

import (
	"fmt"
	"mq-server/config"
	"mq-server/service"
)

func main() {
	config.Init()
	fmt.Println("初始化完成")

	forerver := make(chan bool)

	service.CreateTask()

	<-forerver

}
