package v1

import (
	"github.com/gin-gonic/gin"
	"spiderman-center/common/routes/base"
	"spiderman-center/app"
	"spiderman-center/common/model"
	"fmt"
)

type EtcdRouter struct {
	base.BaseRouter
}

func NewEtcdRouter() *EtcdRouter {
	return new(EtcdRouter)
}

func (r *EtcdRouter) Load(group *gin.RouterGroup) {
	group.GET("/ips", r.ipsHandler)
}

func (r *EtcdRouter) ipsHandler(c *gin.Context) {
	var resp model.IpList
	port := app.GetApp().Config.EtcdServer.Port
	for _, v := range app.GetApp().Config.EtcdServer.Urls{
		ip := fmt.Sprintf("%s:%d", v, port)
		resp.IpList = append(resp.IpList, ip)
	}
	c.JSON(200, resp)
}
