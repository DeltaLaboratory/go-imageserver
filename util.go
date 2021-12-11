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
	_, webpExists := os.Stat("./images/" + hash + ".webp")
	_, avifExists := os.Stat("./images/" + hash + ".avif")
	return (webpExists == nil) && (avifExists == nil)
}

func validate(hash string) bool {
	if len(hash) != 128 {
		return false
	}
	for _, r := range hash {
		if (r < 'a' || r > 'f') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}
