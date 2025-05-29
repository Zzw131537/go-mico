package main

import (
	"fmt"
	"user/config"
)

func main() {
	config.Init()
	fmt.Println("Hello")
}
