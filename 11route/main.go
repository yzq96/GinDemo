package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//加载解析模板
	r.LoadHTMLFiles("templates/404.html", "templates/views/404.html")

	/*普通路由*/
	//定义访问不存在的路由情况
	r.NoRoute(func(c *gin.Context) {
		//test.html 是在模板里define的名称 不定义的话默认使用文件名称
		c.HTML(http.StatusNotFound, "test2.html", nil)
	})
	r.GET("/404", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.html", nil)
	})

	/*路由组*/
	userGroup := r.Group("/user")
	{
		userGroup.GET("/index", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"route": "/user/index"}) })
		userGroup.GET("/yzq", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"route": "/user/yzq"}) })

		//路由组中嵌套路由
		userGroup2 := userGroup.Group("/next")
		{
			userGroup2.GET("/gogogo", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"route": "/user/next/gogogo"}) })
		}
	}
	r.Run(":9000")
}
