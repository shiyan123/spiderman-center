package routes

import (
	"spiderman-center/common/routes/base"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PingRouter struct {
	base.BaseRouter
}

func NewPingRouter() *PingRouter {
	return &PingRouter{}
}

func (s *PingRouter) Load(routerGroup *gin.RouterGroup) {
	routerGroup.HEAD("", pingHandler)
	routerGroup.GET("", pingHandler)
	routerGroup.POST("", pingHandler)
	routerGroup.GET("/customer/ping", pingHandler)
}

func pingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
