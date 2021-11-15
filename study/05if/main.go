package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

//获取当前执行程序的路径
func getCurrentPath() string {
	//if语法糖，和下面if效果一致
	// if ex, err := os.Executable(); err == nil {
	// 	return filepath.Dir(ex)
	// }

	ex, err := os.Executable()
	if err == nil {
		return filepath.Dir(ex)
	}

	return "./"
}
func main() {
	r := gin.Default()

	// gin.H 是map[string]interface{}的缩写
	r.GET("/someJSON", func(c *gin.Context) {
		// 方式一：自己拼接JSON
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})
	r.GET("/moreJSON", func(c *gin.Context) {
		/* // 方法二.1：使用结构体
		var msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}

		msg.Name = "小王子"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.JSON(http.StatusOK, msg) */

		/* //方法二.2：使用map
		m1 := map[string]interface{}{
			"name": "yzq",
			"age":  12,
		}
		c.JSON(http.StatusOK, m1) */

		//方法三.3：使用slice
		s1 := []string{"as", "12"}
		c.JSON(http.StatusOK, s1)
		//结果：
		/*[
		    "as",
		    "12"
		]*/
	})

	r.GET("getCurrentPath", func(c *gin.Context) {
		path := getCurrentPath()
		c.JSON(http.StatusOK, gin.H{"message": path})
	})
	r.Run(":8080")
}
