package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

const Version string = "0.1.0-alpha.1"

var config Config

func main() {
	var err error
	log.Println("Start go-imageserver " + Version)
	log.Println("Loading config...")
	config, err = loadConfig()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("Done. Starting server...")
	gin.SetMode(gin.ReleaseMode)
	server := gin.New()
	if config.MemoryLimit > 0 {
		log.Println("MemoryLimit: " + strconv.Itoa(int(config.MemoryLimit)) + "MiB")
		server.MaxMultipartMemory = config.MemoryLimit << 20
	}
	if config.UploadSizeLimit > 0 {
		log.Println("SizeLimit: " + strconv.Itoa(int(config.UploadSizeLimit)) + "MiB")
		server.Use(sizeLimitMiddleware)
	}
	server.POST("/upload", uploadHandler)
	server.GET("/image/:id", imageHandler)
	server.GET("/favicon.ico", faviconHandler)
	server.NoRoute(noRouteHandler)
	if err := server.Run(":80"); err != nil {
		log.Println(err)
		return
	}
}
