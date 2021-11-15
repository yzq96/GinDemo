package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "order/index"})
}

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "order/hello"})
}
func LoadOrder(e *gin.Engine) {
	e.GET("/index", index)
	e.GET("/hello", hello)
}
