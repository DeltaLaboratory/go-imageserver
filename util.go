package main

import (
	"log"
	"mime/multipart"
	"os"
)

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
