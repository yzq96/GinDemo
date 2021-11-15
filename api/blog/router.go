package blog

import (
	"github.com/gin-gonic/gin"
)

func Routers(e *gin.Engine) {
	e.GET("/api/post", postHandler)
	e.GET("/api/comment", commentHandler)
}
