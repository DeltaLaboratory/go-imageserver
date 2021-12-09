package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	server.POST("/upload", uploadHandler)
	server.GET("/image/:id", imageHandler)
	server.GET("/favicon.ico", faviconHandler)
	server.NoRoute(noRouteHandler)
	if err := server.Run(":80"); err != nil {
		log.Println(err)
		return
	}
}
