package main

import (
	"bytes"
	"errors"
	"github.com/Kagami/go-avif"
	"github.com/chai2010/webp"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
)

func DecodeImage(input []byte) (image.Image, error) {
	var err error
	var img image.Image
	contentType := http.DetectContentType(input)
	switch contentType {
	case "image/jpeg":
		img, err = jpeg.Decode(bytes.NewReader(input))
	case "image/png":
		img, err = png.Decode(bytes.NewReader(input))
	case "image/gif":
		img, err = gif.Decode(bytes.NewReader(input))
	case "image/webp":
		img, err = webp.Decode(bytes.NewReader(input))
	default:
		return nil, errors.New("not supported image type")
	}
	if err != nil {
		return nil, err
	}
	return img, nil
}

func EncodeWebp(input image.Image) (bytes.Buffer, error) {
	var err error
	var buffer bytes.Buffer
	if err = webp.Encode(&buffer, input, &webp.Options{
		Lossless: true,
		Exact:    true,
	}); err != nil {
		return buffer, err
	}
	return buffer, nil
}

func EncodeAvif(input image.Image) (bytes.Buffer, error) {
	var err error
	var buffer bytes.Buffer
	if err = avif.Encode(&buffer, input, &avif.Options{
		Threads: 8,
		Quality: 48,
	}); err != nil {
		return buffer, err
	}
	return buffer, nil
}
