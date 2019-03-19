package main

import (
	"log"
	"spiderman-center/app"

	"spiderman-center/service"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := app.GetApp().Prepare(); err != nil {
		panic(err)
	}
	s, err := service.DisCoveryService()
	if err != nil {
		panic(err)
	}
	go service.MonitorNodes(s)

	server := app.GetApp().Config.Server
	if err := Listen(server.Address, server.Port); err != nil {
		panic(err)
	}
}