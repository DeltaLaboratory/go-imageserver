package main

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"lukechampine.com/blake3"
	"os"
)

func faviconHandler(c *gin.Context) {
	c.File("./statics/favicon.ico")
}

func noRouteHandler(c *gin.Context) {
	c.File("./statics/notfound.webp")
}

// upload single image file
func uploadHandler(c *gin.Context) {
	formFile, err := c.FormFile("image")
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request : no such file",
		})
		return
	}
	buf, err := formFileConvert(formFile)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request : cannot read file",
		})
		return
	}
	hash := blake3.Sum512(buf)
	stringHash := hex.EncodeToString(hash[:])
	if exists(stringHash) == true {
		c.JSON(200, gin.H{
			"message": "success",
			"id":      stringHash,
		})
		return
	}
	decodedImage, err := DecodeImage(buf)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "invalid request : cannot decode image",
		})
		return
	}
	convertedWebp, err := EncodeWebp(decodedImage)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "invalid request : cannot encode image to webp",
		})
		return
	}
	convertedAvif, err := EncodeAvif(decodedImage)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "invalid request : cannot encode image to avif",
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
			"message": "invalid request : cannot write file to storage",
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
			"message": "invalid request : cannot write file to storage",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"id":      stringHash,
	})
	return
}

func imageHandler(c *gin.Context) {
	id := c.Param("id")
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
