package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//定义中间件
func statCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		log.Println(start)
		c.Set("name", "杨振强")                                  //可以通过c.Set在请求
		c.Set("name1", map[string]string{"1": "a", "2": "b"}) //可以通过c.Set在请求
		c.Set("name2", []int{1, 2, 3})                        //可以通过c.Set在请求
		//调用该请求的剩余处理程序
		c.Next()
		//不调用该请求的剩余处理程序
		// c.Abort()
		//计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}

func main() {
	//定义一个加载了这两个全局中间件的Logger(), Recovery()路由
	// r := gin.Default()

	//定义一个不含任何中间件的空路由
	r := gin.New()
	//注册一个全局中间价
	r.Use(statCost())
	r.GET("/test", func(c *gin.Context) {
		name2 := c.MustGet("name2")
		name1 := c.MustGet("name1")
		name := c.MustGet("name").(string) //从上下文取值
		log.Println(name)
		log.Println(name1)
		log.Println(name2)
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/test2", statCost(), func(c *gin.Context) {
		name := c.MustGet("name").(string)
		log.Println(name)
		c.JSON(http.StatusOK, gin.H{"msg": "单路由中间件"})
	})
	//路由组中间件 写法1：
	shopGroup := r.Group("/shop", statCost())
	{
		shopGroup.GET("/test3", func(c *gin.Context) {
			name := c.MustGet("name").(string)
			log.Println(name)
			c.JSON(http.StatusOK, gin.H{"msg": "路由组中间件  写法1"})
		})
	}

	//路由组中间件 写法2：
	userGroup := r.Group("/user")
	userGroup.Use(statCost())
	{
		userGroup.GET("/test4", func(c *gin.Context) {
			name := c.MustGet("name").(string)
			log.Println(name)
			c.JSON(http.StatusOK, gin.H{"msg": "路由组中间件  写法2"})
		})
	}

	/*
		中间件注意事项
		gin默认中间件
		gin.Default()默认使用了Logger和Recovery中间件，其中：

		Logger中间件将日志写入gin.DefaultWriter，即使配置了GIN_MODE=release。
		Recovery中间件会recover任何panic。如果有panic的话，会写入500响应码。
		如果不想使用上面两个默认的中间件，可以使用gin.New()新建一个没有任何默认中间件的路由。

		gin中间件中使用goroutine
		当在中间件或handler中启动新的goroutine时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）。
	*/
	r.Run(":9000")
}
