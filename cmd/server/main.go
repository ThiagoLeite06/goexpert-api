package main

import "github.com/ThiagoLeite06/goexpert/api-9/configs"

func main() {
	config, _ := configs.LoadConfig("./cmd/server/")
	println(config.DBDriver)
}
