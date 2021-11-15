package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func test1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "/user/test1"})
}

func test2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "/user/test2"})
}
func LoadUser(e *gin.Engine) {
	userGroup := e.Group("/user")
	userGroup.GET("/test1", test1)
	userGroup.GET("/test2", test2)
}
