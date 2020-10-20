package main

import (
	"myserver/app/server"
)

func main() {
	g := server.NewServer()
	g.Start()
}
