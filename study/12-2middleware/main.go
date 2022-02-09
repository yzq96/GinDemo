package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func m1(c *gin.Context) {
	start := time.Now()
	fmt.Println("m1 start")
	c.Next()  //开始此后的程序
	c.Abort() //终止程序
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 end")
}
func indexHandleFunc(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "message"})
}
func main() {
	r := gin.Default()
	r.Use(m1)
	userGroup := r.Group("/user")
	userGroup.GET("/index", m1, indexHandleFunc)
	r.Run()
}
