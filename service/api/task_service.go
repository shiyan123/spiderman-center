package api

import (
	"fmt"
	"spiderman-center/common/model"
	"spiderman-center/service"
	"sync"
)

type TaskService struct {
}

var (
	serviceOnce sync.Once
	serviceT    *TaskService
)

func GetTaskService() *TaskService {
	serviceOnce.Do(func() {
		serviceT := &TaskService{}
		serviceT.init()
	})
	return serviceT
}

func (s *TaskService) init() {

}

func (s *TaskService) Scheduling(t *model.TaskInfo) {

	var nodes []*model.Node
	group := service.GetService().RunningGroup
	for k, v := range group {
		if k == t.GroupName {
			fmt.Sprintf("分组名:%s  节点数量:%d \n", k, len(v))
			nodes = v
			break
		}
	}

	for _, v := range nodes {
		if _, ok := v.TaskMap[t.TaskId]; ok {
			fmt.Println("任务已经存在")
			return
		} else {
			//todo 开始调度
			if !s.scheduling(v, t) {
				continue
			}
			//调度完成并发送
			err := service.GetService().PutNode(t.GroupName, v.ID, v)
			if err != nil {
				return
			}
			return
		}
	}
	return
}

func (s *TaskService) scheduling(node *model.Node, t *model.TaskInfo) bool {
	return true
}