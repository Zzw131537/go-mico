package weblib

import (
	"api-gateway/weblib/handlers"
	"api-gateway/weblib/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRoute(service ...interface{}) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.Use(middleware.Cors(), middleware.InitMiddlewares(service))

	store := cookie.NewStore([]byte("something-very-secret"))
	ginRouter.Use(sessions.Sessions("mysession", store))

	v1 := ginRouter.Group("api/v1")
	{

		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(200, "success")
		})

		// 用户服务
		v1.POST("/user/register", handlers.UserRegister)
		v1.POST("/user/login", handlers.UserLogin)
	}
	return ginRouter

}
