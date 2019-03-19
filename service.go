package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"spiderman-center/routes"
)

func Listen(address string, port int) error {
	r := gin.Default()

	routes.Load(r)
	return r.Run(fmt.Sprintf("%s:%d", address, port))
}