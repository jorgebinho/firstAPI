package main

import "github.com/jorgebinho/firstApi/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
