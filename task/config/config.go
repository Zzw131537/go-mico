package config

import (
	"fmt"
	"strings"
	"task/model"

	"gopkg.in/ini.v1"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassWord string
	RabbitMQHost     string
	RabbitMQPort     string
)

func Init() {
	file, err := ini.Load("./config/config.ini")
	if err != nil {
		fmt.Println("配置文件读取错误，请检测文件路径:", err)
	}
	LoadMySqlData(file)
	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true"}, "")
	model.Database(path)

	LoadRabbitMQData(file)

	// 连接RabbitMQ
	pathRabbitMQ := strings.Join([]string{RabbitMQ, "://", RabbitMQUser, ":", RabbitMQPassWord, "@", RabbitMQHost, ":", RabbitMQPort, "/"}, "")

	model.RabbitMQ(pathRabbitMQ)

}

func LoadRabbitMQData(file *ini.File) {
	RabbitMQ = file.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassWord = file.Section("rabbitmq").Key("RabbitMQPassWord").String()
	RabbitMQHost = file.Section("rabbitmq").Key("RabbitMQHost").String()
	RabbitMQPort = file.Section("rabbitmq").Key("RabbitMQPort").String()
}

func LoadMySqlData(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()

}
