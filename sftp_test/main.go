package main

import (
	"net/http"
	"sftp_test/controllers"

	"github.com/gin-gonic/gin"
)

var C *gin.Context

func main() {

	// client := config.Sftp_connection()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello"})
	})
	r.GET("/files", controllers.Listcontroller)
	r.Run(":8081")
}
