package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/service"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
)

func UserRegister(ctx *gin.Context) {
	var userReq service.UserRequest

	PanicIfUserError(ctx.Bind(&userReq))

	// 从 gin.Key中取出服务实例
	userService := ctx.Keys["userServices"].(service.UserService)

	userResp, err := userService.UserRegister(context.Background(), &userReq)

	fmt.Println(userResp)
	PanicIfUserError(err)
	ctx.JSON(200, gin.H{
		"data": userResp,
	})

}

// 用户登录
func UserLogin(ctx *gin.Context) {
	var userReq service.UserRequest

	if err := ctx.Bind(&userReq); err != nil {
		logger.Info(err.Error())
	}

	// 从 gin.Key中取出服务实例
	userService := ctx.Keys["userService"].(service.UserService)

	userResp, err := userService.UserLogin(context.Background(), &userReq)
	PanicIfUserError(err)

	token, err := utils.GenerateToken(uint(userResp.UserDetail.ID))
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"user":  userResp.UserDetail,
			"token": token,
		},
	})

}
