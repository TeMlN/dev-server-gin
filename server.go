package main

import (
	"gin-server/config"
	"gin_server/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config.StartConfig()
	database.Connect()
	r := gin.Default() //router 생성

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:3030 (for windows "localhost:3030")

}
