package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

/*
重定向
*/
func main() {
	route := gin.Default()

	//http重定向
	route.GET("/http/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	//路由重定向
	route.GET("/route/redirect", func(c *gin.Context) {
		c.Request.URL.Path = "/route/redirect2"
		route.HandleContext(c)
	})
	route.GET("/route/redirect2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	route.Run(":9000")
}
