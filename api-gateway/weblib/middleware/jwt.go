package middleware

import (
	"api-gateway/pkg/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}
		code = 200

		token := c.GetHeader("Authorization")

		if len(token) > 7 && strings.HasPrefix(strings.ToUpper(token), "BEARER ") {
			token = token[7:]
		}

		if token == "" {
			code = 404
		} else {
			claims, err := utils.ParseToken(token)
			if err != nil {
				code = 500
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = 500
			}
		}
		if code != 200 {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    "",
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
