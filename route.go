package main

import (
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"lukechampine.com/blake3"
	"mime/multipart"
	"os"
)

func faviconHandler(c *gin.Context) {
	c.File("./statics/favicon.ico")
}

func formFileConvert(form *multipart.FileHeader) ([]byte, error) {
	buf := make([]byte, form.Size)
	file, err := form.Open()
	if err != nil {
		return nil, err
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	}(file)
	if _, err = file.Read(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

func exists(hash string) bool {
	_, err := os.Stat("./images/" + hash + ".webp")
	return err == nil
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
