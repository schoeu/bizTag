package main

import (
	"config"
	"server"
)

func main() {
	config.ReadConf()
	server.RunServer()
	server.CloseDB()
}
