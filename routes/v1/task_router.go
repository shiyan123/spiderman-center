package v1

import (
	"github.com/gin-gonic/gin"
	"spiderman-center/common/routes/base"
	"spiderman-center/common/model"
	"github.com/shiyan123/marvel.sy/common/rd"

	"github.com/shiyan123/marvel.sy/common/errors"
	"net/http"
	"spiderman-center/service/api"
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
	var req model.TaskInfo
	if err := c.ShouldBindJSON(&req);err !=nil{
		c.JSON(http.StatusOK, errors.ErrInvalidParams)
		return
	}
	api.GetTaskService().Scheduling(&req)
	c.JSON(http.StatusOK, rd.Data("success"))
}