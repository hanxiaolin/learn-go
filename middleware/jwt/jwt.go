package jwt

import (
	"net/http"
	"time"

	"code.shihuo.cn/gin-demo/pkg/e"
	"code.shihuo.cn/gin-demo/pkg/util"

	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code e.ResponseType
		var data interface{}

		// todo
		if c.Request.Method == "GET" || c.Request.Method == "POST" {
			c.Next()
			return
		}

		code = e.SUCCESS
		token := c.Query("token")
		if token == "" {
			code = e.ERROR_AUTH

		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
