package handlers

import (
	"api-gateway/pkg/utils"
	"api-gateway/service"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(ctx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))

	userService := ctx.Keys["userService"].(service.UserService)

	userResp, err := userService.UserRegister(context.Background(), &userReq)

	PanicIfUserError(err)

	ctx.JSON(http.StatusOK, gin.H{
		"data": userResp,
	})

}

func UserLogin(ctx *gin.Context) {
	var userReq service.UserRequest
	PanicIfUserError(ctx.Bind(&userReq))

	userService := ctx.Keys["userService"].(service.UserService)

	userResp, err := userService.UserLogin(context.Background(), &userReq)

	PanicIfUserError(err)

	fmt.Println("用户Id为: " + string(userResp.UserDetail.ID))
	token, err := utils.GenerateToken(uint(userResp.UserDetail.ID))
	ctx.JSON(http.StatusOK, gin.H{
		"code": userResp.Code,
		"data": gin.H{
			"resp":  userResp.UserDetail,
			"token": token,
		},
	})
}
