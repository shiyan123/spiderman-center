package v1

import (
	"github.com/gin-gonic/gin"
	"spiderman-center/common/routes/base"
)

type TaskRouter struct {
	base.BaseRouter
}

func NewTaskRouter() *TaskRouter {
	return new(TaskRouter)
}

func (r *TaskRouter) Load(group *gin.RouterGroup) {
	group.POST("", r.addHandler)
}

func (r *TaskRouter) addHandler(c *gin.Context) {

	c.JSON(200, "test")
}