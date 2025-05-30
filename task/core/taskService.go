package core

import (
	"context"
	"encoding/json"
	"errors"
	"task/model"
	"task/service"

	"github.com/streadway/amqp"
)

// 将信息放到消息队列中,生产者
func (*TaskService) CreateTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	// 进行连接
	ch, err := model.MQ.Channel()
	if err != nil {
		err = errors.New("rabbitMQ channel err:" + err.Error())
		return err
	}

	q, _ := ch.QueueDeclare("task_queue", true, false, false, false, nil)

	body, _ := json.Marshal(req)

	// 发布到队列中
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         body,
	})

	if err != nil {
		err = errors.New("rabbitMQ channel err:" + err.Error())
		return err
	}
	return nil
}

// 获取备忘录列表
func (*TaskService) GetTaskList(ctx context.Context, req *service.TaskRequest, resp *service.TaskListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}

	var taskData []model.Task
	var count int64

	if err := model.DB.Model(&model.Task{}).Offset(int(req.Start)).Limit(int(req.Limit)).Where("uid=?", req.Uid).Find(&taskData).Error; err != nil {
		return errors.New("mysql Not find" + err.Error())
	}

	model.DB.Model(&model.Task{}).Where("uid=?", req.Uid).Count(&count)

	var taskRes []*service.TaskModel

	for _, item := range taskData {
		taskRes = append(taskRes, BuildTask(&item))
	}

	resp.TaskList = taskRes
	resp.Count = uint32(count)
	return nil
}

// 获取备忘录详细信息
