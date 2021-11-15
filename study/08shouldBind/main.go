package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type login struct {
	User     string `json:"user" form:"user" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

/*
ShouldBind会按照下面的顺序解析请求中的数据完成绑定：

如果是 GET 请求，只使用 Form 绑定引擎（query）。
如果是 POST 请求，首先检查 content-type 是否为 JSON 或 XML，然后再使用 Form（form-data）。
*/
func main() {
	// slice1 := []string{"1", "2", "3"}
	// s, err := json.Marshal(slice1)
	// fmt.Printf("%#v\n", slice1)
	// fmt.Printf("%#v\n", string(s))
	// fmt.Printf("%#v\n", err)

	// map1 := map[string]string{"User": "a", "Password": "b"}
	// m1, err := json.Marshal(map1)
	// fmt.Printf("%#v\n", slice1)
	// fmt.Printf("%#v\n", string(m1))
	// fmt.Printf("%#v\n", err)

	// login1 := new(login)
	// data := []byte(m1)
	// err = json.Unmarshal(data, &login1)
	// if err != nil {
	// 	fmt.Printf("%#v\n", err)
	// }
	// fmt.Printf("%#v\n", login1)
	// return
	router := gin.Default()

	//
	router.POST("/loginJSON", func(c *gin.Context) {
		var login login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

	})

	router.POST("loginForm", func(c *gin.Context) {
		var login login
		//ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("%v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.GET("loginQuery", func(c *gin.Context) {
		var login login
		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("%v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
	router.Run(":9000")
}
