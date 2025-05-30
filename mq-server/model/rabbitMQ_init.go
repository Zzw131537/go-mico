package model

import (
	"fmt"

	"github.com/streadway/amqp"
)

var MQ *amqp.Connection

func RabbitMQ(connString string) {
	conn, err := amqp.Dial(connString)

	if err != nil {
		fmt.Println("RabbitMQ 连接错误!")
		panic(err)
	}
	MQ = conn
}
