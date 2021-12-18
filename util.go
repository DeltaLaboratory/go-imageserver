package main

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"runtime"
)

type WebpConfig struct {
	Lossless bool    `yaml:"Lossless"`
	Quality  float32 `yaml:"Quality"`
	Exact    bool    `yaml:"Exact"`
}

type AvifConfig struct {
	Threads int `yaml:"Threads"`
	Speed   int `yaml:"Speed"`
	Quality int `yaml:"Quality"`
}

type Config struct {
	WebpOption WebpConfig `yaml:"WebpOption"`
	AvifOption AvifConfig `yaml:"AvifOption"`
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
	_, webpExists := os.Stat("./images/" + hash + ".webp")
	_, avifExists := os.Stat("./images/" + hash + ".avif")
	return (webpExists == nil) && (avifExists == nil)
}

func loadConfig() (Config, error) {
	var config Config
	_, err := os.Stat("./config/config.yml")
	if err != nil {
		data, err := yaml.Marshal(Config{
			WebpOption: WebpConfig{
				Lossless: true,
				Quality:  100,
				Exact:    false,
			},
			AvifOption: AvifConfig{
				Threads: runtime.NumCPU(),
				Speed:   1,
				Quality: 100,
			},
		})
		if err != nil {
			log.Println(err)
			return config, err
		}
		if err = ioutil.WriteFile("./config/config.yml", data, 0644); err != nil {
			log.Println(err)
			return config, err
		}
	}
	data, err := ioutil.ReadFile("./config/config.yml")
	if err != nil {
		log.Println(err)
		return config, err
	}
	if err = yaml.Unmarshal(data, &config); err != nil {
		log.Println(err)
		return config, err
	}
	if !validateConfig(config) {
		return config, errors.New("invalid config")
	}
	return config, nil
}

func validateConfig(in Config) bool {
	if in.WebpOption.Quality < 0 || in.WebpOption.Quality > 100 {
		return false
	}
	if in.AvifOption.Quality < 0 || in.AvifOption.Quality > 63 {
		return false
	}
	if in.AvifOption.Speed < 0 || in.AvifOption.Speed > 8 {
		return false
	}
	return true
}

func validate(hash string) bool {
	if len(hash) != 128 {
		return false
	}
	for _, r := range hash {
		if (r < 'a' || r > 'f') && (r < '0' || r > '9') { // [a-f && 0-9]
			return false
		}
	}
	return true
}
