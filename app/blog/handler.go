package blog

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func postHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "blog/post"})
}

func commentHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "blog/comment"})
}
