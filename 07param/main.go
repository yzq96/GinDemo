package main

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

//获取querystring参数
func getQuery(c *gin.Context) {
	username := c.DefaultQuery("username", "杨振强") //没有的话取默认值
	message := c.Query("message")
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"message":  message,
	})
}

//获取post参数（form-data x-www-form-urlencoded）
func getForm(c *gin.Context) {
	username := c.PostForm("username")
	message := c.DefaultPostForm("message", "无")
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"message":  message,
	})
}

//获取post参数中raw(json)方式的参数
func getJson(c *gin.Context) {
	b, _ := c.GetRawData()
	var m map[string]interface{}
	_ = json.Unmarshal(b, &m)
	c.JSON(http.StatusOK, m)
}
func getPath(c *gin.Context) {
	username := c.Param("username")
	msg := c.Param("msg")
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"message":  msg,
	})
}
func main() {
	r := gin.Default() //Default返回一个已经附加了Logger和Recovery中间件的Engine实例。
	r.GET("user/search", getQuery)
	r.POST("user/search", getForm)
	r.POST("user/json", getJson)
	r.POST("user/path/:username/:msg", getPath)
	r.Run(":9000")
}
