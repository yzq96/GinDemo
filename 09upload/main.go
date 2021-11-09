package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLFiles("templates/upload.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{
			"title": "upload",
		})
	})
	r.LoadHTMLFiles("templates/uploads.html")
	r.GET("/indexs", func(c *gin.Context) {
		c.HTML(http.StatusOK, "uploads.html", gin.H{
			"title": "uploads",
		})
	})

	//单文件生上传
	r.POST("/upload", func(c *gin.Context) {
		file, err := c.FormFile("f1")
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"message": err.Error(),
			})
			return
		}
		log.Println(file.Filename)
		dst := fmt.Sprintf("C:/tmp/%s", file.Filename)
		fmt.Printf("%#v\n", dst)
		// 上传文件到指定的目录
		c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})

	//多文件上传
	r.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["file"]
		for index, file := range files {
			log.Println(file.Filename)
			dst := fmt.Sprintf("C:/tmp/%s_%d", file.Filename, index)
			//上传文件到指定目录
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})
	r.Run(":9000")
}
