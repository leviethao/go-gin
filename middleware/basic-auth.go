package middleware

import (
	"github.com/gin-gonic/gin"
)

func BasicAuth() gin.HandlerFunc {
	return gin.BasicAuth(gin.Accounts{
		"haolv": "haolv123456",
	})
}
