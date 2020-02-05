package main

import (
	pingcontroller "github.com/aandresvelasquezn/MarvelComics/internal/Controller/Ping"
	"github.com/gin-gonic/gin"
)

func main() {
	//dotenv := tools.GetDotEnvVariable("PUBLIC_KEY")
	//fmt.Println(dotenv)
	//tools.GetComics()
	r := gin.Default()
	r.GET("/ping", pingcontroller.Ping)
	r.Run()
}
