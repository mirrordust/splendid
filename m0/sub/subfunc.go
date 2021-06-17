package subf

import (
	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"123🐸": "456🐻",
		})
	})
	return r
}
