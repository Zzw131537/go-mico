package core

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"task/model"
	"task/service"

	"github.com/streadway/amqp"
	"gorm.io/gorm"
)

// 消息通过MQ进行传递

func (task *TaskService) CreateTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	ch, err := model.MQ.Channel()

	fmt.Println(req.Content, req.Title, req.Status)

	if err != nil {
		err = errors.New("rabbit Mq channel error: " + err.Error())
	}
	q, _ := ch.QueueDeclare("task_queue", true, false, false, false, nil)

	body, _ := json.Marshal(req) // 请求序列化
	err = ch.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		Body:         body,
	})
	if err != nil {
		err = errors.New("rabbit Mq channel error: " + err.Error())
	}
	return nil

}

// 获取备忘录列表
func (*TaskService) GetTasksList(ctx context.Context, req *service.TaskRequest, resp *service.TaskListResponse) error {
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
func (*TaskService) GetTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	taskData := model.Task{}

	err := model.DB.Model(&model.Task{}).Where("id = ?", req.Id).First(&taskData).Error
	if err != nil {
		fmt.Println("查询数据库错误!")
		return err
	}
	taskRes := BuildTask(&taskData)
	resp.TaskDetail = taskRes
	return nil
}

// 修改信息
// 修改信息
func (*TaskService) UpdateTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	taskModel := model.Task{}

	fmt.Printf("=== 更新任务调试 ===\n")
	fmt.Printf("请求参数: ID=%d, UID=%d, Title=%s, Content=%s, Status=%d\n",
		req.Id, req.Uid, req.Title, req.Content, req.Status)

	// 1. 先查找任务
	err := model.DB.Where("id = ? and uid = ?", req.Id, req.Uid).First(&taskModel).Error
	if err != nil {
		fmt.Printf("查找任务失败: ID=%d, UID=%d, 错误: %v\n", req.Id, req.Uid, err)
		return err
	}

	fmt.Printf("找到任务: ID=%d, 原标题=%s\n", taskModel.ID, taskModel.Title)

	// 2. 更新字段
	taskModel.Title = req.Title
	taskModel.Content = req.Content
	taskModel.Status = req.Status

	fmt.Printf("更新后任务: Title=%s, Content=%s, Status=%d\n",
		taskModel.Title, taskModel.Content, taskModel.Status)

	// 3. 保存 - 关键：传递指针 &
	err = model.DB.Save(&taskModel).Error
	if err != nil {
		fmt.Printf("保存失败: %v\n", err)
		return errors.New("保存数据错误: " + err.Error())
	}

	fmt.Println("更新成功")
	return nil
}

// 删除备忘录
func (*TaskService) DeleteTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
	taskModel := model.Task{}

	fmt.Printf("=== 删除任务调试 ===\n")
	fmt.Printf("请求参数: ID=%d, UID=%d\n", req.Id, req.Uid)

	// 1. 先查找任务是否存在
	err := model.DB.Where("id = ? and uid = ?", req.Id, req.Uid).First(&taskModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("任务不存在: ID=%d, UID=%d\n", req.Id, req.Uid)
			return fmt.Errorf("任务不存在")
		}
		fmt.Printf("查找任务失败: %v\n", err)
		return fmt.Errorf("数据库错误: %v", err)
	}

	fmt.Printf("找到要删除的任务: ID=%d, Title=%s\n", taskModel.ID, taskModel.Title)

	// 2. 执行删除
	result := model.DB.Delete(&taskModel)
	if result.Error != nil {
		fmt.Printf("删除失败: %v\n", result.Error)
		return fmt.Errorf("删除失败: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		fmt.Println("没有记录被删除")
		return fmt.Errorf("删除失败，记录不存在")
	}

	fmt.Printf("删除成功，影响行数: %d\n", result.RowsAffected)
	return nil
}
