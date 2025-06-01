package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/service"
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTasksList(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))

	taskService := ctx.Keys["taskService"].(service.TaskService)

	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))

	taskReq.Uid = uint64(claim.Id)

	// 调用服务端函数

	taskResp, err := taskService.GetTasksList(context.Background(), &taskReq)

	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"code":  "200",
		"tasks": taskResp.TaskList,
		"count": taskResp.Count,
	})
}

func CreateTask(ctx *gin.Context) {

	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))

	taskService := ctx.Keys["taskService"].(service.TaskService)

	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))

	taskReq.Uid = uint64(claim.Id)

	// 调用服务端函数

	taskResp, err := taskService.CreateTask(context.Background(), &taskReq)

	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"code": "200",
		"date": taskResp.TaskDetail,
	})

}

func GetTask(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))

	taskService := ctx.Keys["taskService"].(service.TaskService)

	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))

	taskReq.Uid = uint64(claim.Id)

	id, _ := strconv.Atoi(ctx.Param("id")) // 获取 task_id
	taskReq.Id = uint64(id)
	// 调用服务端函数

	taskResp, err := taskService.GetTask(context.Background(), &taskReq)

	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"code": "200",
		"date": taskResp.TaskDetail,
	})

}
func UpdateTask(ctx *gin.Context) {
	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))

	taskService := ctx.Keys["taskService"].(service.TaskService)

	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))

	taskReq.Uid = uint64(claim.Id)

	id, _ := strconv.Atoi(ctx.Param("id")) // 获取 task_id
	taskReq.Id = uint64(id)
	// 调用服务端函数

	taskResp, err := taskService.UpdateTask(context.Background(), &taskReq)

	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"code": "200",
		"date": taskResp.TaskDetail,
	})
}

func DeleteTask(ctx *gin.Context) {

	var taskReq service.TaskRequest
	PanicIfTaskError(ctx.Bind(&taskReq))

	taskService := ctx.Keys["taskService"].(service.TaskService)

	claim, _ := utils.ParseToken(ctx.GetHeader("Authorization"))

	taskReq.Uid = uint64(claim.Id)

	id, _ := strconv.Atoi(ctx.Param("id")) // 获取 task_id
	taskReq.Id = uint64(id)
	// 调用服务端函数

	taskResp, err := taskService.DeleteTask(context.Background(), &taskReq)

	PanicIfTaskError(err)
	ctx.JSON(200, gin.H{
		"code": "200",
		"date": taskResp.TaskDetail,
	})

}
