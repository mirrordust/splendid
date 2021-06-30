package web

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/mirrordust/splendour/m0/repo"
)

func Server() *gin.Engine {
	var tags []repo.Tag
	err := repo.FindAll(&tags, "1=1")
	if err != nil {
		log.Panicln("DB error")
	}

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, tags)
	})
	return r
}
