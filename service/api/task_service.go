package api

import (
	"sync"
	"spiderman-center/common/model"
	"spiderman-center/service"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"context"
	"encoding/json"
)

type TaskService struct {
}

var (
	serviceOnce sync.Once
	serviceT     *TaskService
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

func (s *TaskService) SendTask(t *model.TaskInfo){

	group := service.GetService().RunningGroup
	//todo 调度算法
	id := ""
	var node *model.Node
	for k, v := range group{
		if k == t.GroupName {
			for m := range v {
				group[k][m].TaskMap[t.TaskName] = t
				node = group[k][m]
				id = group[k][m].ID
				break
			}
		}
	}

	key := fmt.Sprintf("services/%s/%s", t.GroupName, id)
	re, err := service.GetService().Client.Grant(context.TODO(), 5)
	if err != nil {
		return
	}
	body, _ := json.Marshal(node)
	service.GetService().Client.Put(context.Background(), key, string(body), clientv3.WithLease(re.ID))
}