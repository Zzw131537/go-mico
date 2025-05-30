package middleware

import (
	"github.com/gin-gonic/gin"
)

// 接收服务实例存在 ctx。Key中
func InitMiddlewares(service []interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 将实例存在 gin.Key中
		ctx.Keys = make(map[string]any)
		ctx.Keys["userServices"] = service[0]
		ctx.Next()
	}
}

// // 错误处理中间件
// func ErrorMiddleware() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		defer func() {
// 			if r := recover(); r != nil {
// 				ctx.JSON(200, gin.H{
// 					"code":  "404",
// 					"Hello": "620",
// 					"msg":   fmt.Sprintf("%s", r),
// 				})
// 				ctx.Abort()
// 			}
// 		}()
// 		ctx.Next()
// 	}
// }
