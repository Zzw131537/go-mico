package core

import (
	"task/model"
	"task/service"
)

func BuildTask(item *model.Task) *service.TaskModel {
	taskModel := service.TaskModel{
		Id:         uint64(item.ID),
		Uid:        uint64(item.Uid),
		Title:      item.Title,
		Content:    item.Content,
		StartTime:  item.StartTime,
		EndTime:    item.EndTime,
		Status:     item.Status,
		CreateTime: item.CreatedAt.Unix(),
		UpdateTime: item.UpdatedAt.Unix(),
	}
	return &taskModel
}
