package main

import (
	"github.com/sandeepravi/muninndb/server"
)

func main() {
	server.Setup()
	server.SetupCommands()
	server.Start()
}
