package api

import "sync"

type TaskService struct {
}

var (
	serviceOnce sync.Once
	service     *TaskService
)

func GetTaskService() *TaskService {
	serviceOnce.Do(func() {
		service := &TaskService{}
		service.init()
	})
	return service
}

func (s *TaskService) init() {

}

