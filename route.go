package main

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"lukechampine.com/blake3"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// [GET] /favicon.ico
func faviconHandler(c *gin.Context) {
	c.File("./statics/favicon.ico")
}

// [GET] {noRoute}
func noRouteHandler(c *gin.Context) {
	c.File("./statics/notfound.webp")
}

// [POST] /upload
func uploadHandler(c *gin.Context) {
	content, err := c.GetRawData()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request : no such file",
		})
		return
	}
	hash := blake3.Sum512(content)
	stringHash := hex.EncodeToString(hash[:])
	if exists(stringHash) == true {
		c.JSON(200, gin.H{
			"message": "success",
			"id":      stringHash,
		})
		return
	}
	decodedImage, err := DecodeImage(content)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "cannot decode image",
		})
		return
	}
	convertedWebp, err := EncodeWebp(decodedImage)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot encode image to webp",
		})
		return
	}
	convertedAvif, err := EncodeAvif(decodedImage)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "cannot encode image to avif",
		})
		return
	}
	fileHandler, err := os.Create("./images/" + stringHash + ".webp")
	defer func(fileHandler *os.File) {
		err := fileHandler.Close()
		if err != nil {
			log.Println(err)
		}
	}(fileHandler)
	_, err = fileHandler.Write(convertedWebp.Bytes())
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to write file to storage",
		})
		return
	}
	fileHandler, err = os.Create("./images/" + stringHash + ".avif")
	defer func(fileHandler *os.File) {
		err := fileHandler.Close()
		if err != nil {
			log.Println(err)
		}
	}(fileHandler)
	_, err = fileHandler.Write(convertedAvif.Bytes())
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to write file to storage",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"id":      stringHash,
	})
	return
}

// [GET] /image/:id
func imageHandler(c *gin.Context) {
	id := strings.ToLower(c.Param("id"))
	if validate(id) != true {
		c.JSON(400, gin.H{
			"message": "invalid request : cannot find image",
		})
		return
	}
	format := c.Query("format")
	if format == "" {
		format = "webp"
	}
	if exists(id) == false {
		c.JSON(404, gin.H{
			"message": "not found",
		})
		return
	} else {
		c.File("./images/" + id + "." + format)
		return
	}
}

func sizeLimitMiddleware(c *gin.Context) {
	var w http.ResponseWriter = c.Writer
	fileSize, err := strconv.Atoi(c.Request.Header.Get("Content-Length"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request : invalid Content-Length header",
		})
		c.Abort()
		return
	}
	if int64(fileSize) > config.UploadSizeLimit<<20 {
		c.JSON(400, gin.H{
			"message": "invalid request : file size limit exceeded",
		})
		c.Abort()
		return
	}
	c.Request.Body = http.MaxBytesReader(w, c.Request.Body, config.UploadSizeLimit<<20) // for security reason
	c.Next()
}
