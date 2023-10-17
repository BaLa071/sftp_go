package controllers

import (
	"sftp_test/config"
	"sftp_test/service"

	"github.com/gin-gonic/gin"
)

func Listcontroller(c *gin.Context) {
	service.List(c, config.Sftp_connection())
}
