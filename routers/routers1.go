package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello yzq!",
	})
}

// SetupRouter 配置路由信息
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", helloHandler)
	return r
}
