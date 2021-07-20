package main

import "github.com/gin-gonic/gin"

func InitHttpServerrouter() *gin.Engine {

	gin.SetMode(gin.DebugMode)
	server := gin.Default()

	server.GET("/downloadcfg/:cfg_name", handleDownload)

	server.POST("/uploadcfg", handleUpload)

	return server
}
