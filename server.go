package main

import (
	"gin-server/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //router 생성
	db.Connect()



	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:3030 (for windows "localhost:3030")

}