package shop

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func goodsHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "shop/goods"})
}

func checkoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "shop/checkout"})
}
