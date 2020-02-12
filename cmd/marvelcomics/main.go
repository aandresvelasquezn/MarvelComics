package main

import (
	"net/http"

	pingcontroller "github.com/aandresvelasquezn/MarvelComics/internal/Controller/Ping"
	"github.com/aandresvelasquezn/MarvelComics/tools"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {
	//tools.Comics()
	tools.ComicByID(376)
	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()
	initRoutes()
	router.Run(":8081")
}

func initRoutes() {
	router.GET("/ping", pingcontroller.Ping)
}

// Test edd
func Test() {
	router.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		//value, ok := db[user]
		c.JSON(http.StatusOK, gin.H{"user": user, "value": "1"})

	})
}
