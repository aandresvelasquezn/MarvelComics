package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//dotenv := tools.GetDotEnvVariable("PUBLIC_KEY")
	//fmt.Println(dotenv)
	//tools.GetComics()
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
