package main

import (
	"log"
	"spiderman-center/app"

	"errors"
	"spiderman-center/service"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := app.GetApp().Prepare(); err != nil {
		panic(err)
	}
	s := service.GetService()
	if s == nil {
		panic(errors.New("service is error"))
	}
	go service.MonitorNodes(s)

	server := app.GetApp().Config.Server
	if err := Listen(server.Address, server.Port); err != nil {
		panic(err)
	}
}
