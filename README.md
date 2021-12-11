# go-imageserver
go-imageserver is an image server which automatically optimize non webp and avif images to webp and avif images.

## how to use
### run directly
First, install libaom and gcc as compile dependencies by `sudo apt update && sudo apt install -y libaom-dev build-essential`(debian-based linux) or `sudo yum update && sudo yum install -y libaom-devel gcc`(centos-based linux)\.
Next, clone this repo, and execute the program by `go run main.go`.
### run through docker
`docker pull ghcr.io/deltalaboratory/go-imageserver:latest` and `docker run [some image dir or volume]:/go/src/app/images/ -p 8000:80 ghcr.io/deltalaboratory/go-imageserver:latest`

## supported format
- upload format
  - png
  - jpeg
  - webp
  - gif [partly supported]
- download format
  - webp
  - avif