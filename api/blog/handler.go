package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "api/blog/post"})
}

func commentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "api/blog/comment"})
}
