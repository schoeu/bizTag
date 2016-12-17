package main

import (
	"server"
	"config"
)

func main() {
	config.ReadConf()
	server.RunServer(":8911")

}
