package main

import (
	"flag"
	"log"

	"ds0nt.com/config"

	"ds0nt.com/client"
	"ds0nt.com/server"
)

var (
	configFile = flag.String("conf", "config.toml", "path to toml config")
)

func main() {
	flag.Parse()
	config.Load(*configFile)

	err := client.RenderIndex()
	if err != nil {
		log.Fatal("Render Index Error", err)
	}
	server.Run()
}
