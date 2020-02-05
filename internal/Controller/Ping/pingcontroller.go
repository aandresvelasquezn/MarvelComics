package pingcontroller

import "github.com/gin-gonic/gin"

//Ping return Json
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
