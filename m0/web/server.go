package web

import (
	"fmt"
	"github.com/gin-gonic/gin"

	"github.com/mirrordust/splendour/m0/repo"
)

func Server() *gin.Engine {
	birthdays := db.All()
	// r2 := make(map[string]string)
	// for _, birthday := range birthdays {
	// 	r2[birthday.Name] = birthday.Born.Format("January 2, 2006")
	// }
	fmt.Println(birthdays)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, birthdays)
	})
	return r
}

// func Server() *gin.Engine {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})
// 	return r
// }
