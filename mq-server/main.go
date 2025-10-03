package main

import (
	"mq-server/config"
	"mq-server/service"
)

func main() {
	config.Init()

	forever := make(chan bool)

	service.CreateTask()
	<-forever
}
