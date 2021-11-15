package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// gin.H 是map[string]interface{}的缩写
	r.GET("/someXML", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.XML(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/moreXML", func(c *gin.Context) {
		// 方法二：使用结构体
		type MessageRecord struct {
			Name    string
			Message string
			Age     int
		}
		var msg MessageRecord
		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.XML(http.StatusOK, msg)
	})
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"message": "ok"})
	})
	//有些问题，先搁置以后再说
	r.GET("/someProtoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// label := "test"
		// data := &protoexample.Test{
		// 	Label: &label,
		// 	Reps:  reps,
		// }
		c.ProtoBuf(http.StatusOK, reps)
	})
	r.Run(":8080")
}
