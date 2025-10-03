package wrappers

import (
	"api-gateway/service"
	"strconv"

	"github.com/micro/go-micro/v2/client"
)

func NewTask(id uint64, name string) *service.TaskModel {
	return &service.TaskModel{
		Id:         id,
		Title:      name,
		Content:    "响应超时",
		StartTime:  1000,
		EndTime:    1000,
		Status:     0,
		CreateTime: 1000,
		UpdateTime: 1000,
	}
}

// 降级函数
func DefaultTasks(resp interface{}) {
	models := make([]*service.TaskModel, 0)

	var i uint64

	for i = 0; i < 10; i++ {
		models = append(models, NewTask(i, "降级备忘录"+strconv.Itoa(20+int(i))))
	}

	res := resp.(*service.TaskListResponse)
	res.TaskList = models
}

type taskWrapper struct {
	client.Client
}

func NewTaskWrapper(c client.Client) client.Client {

	return &taskWrapper{c}
}
