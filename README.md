# go-imageserver
go-imageserver is an image server which automatically optimize non webp and avif images to webp and avif images.

## workflows
| branch  | CodeQL                                                                                                                                                                                                                 | Build & Test                                                                                                                                                                                       | Docker Build & Push                                                                                                                                                                                                      |
|---------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| release | [![CodeQL](https://github.com/DeltaLaboratory/go-imageserver/actions/workflows/codeql-analysis.yml/badge.svg?branch=release)](https://github.com/DeltaLaboratory/go-imageserver/actions/workflows/codeql-analysis.yml) | [![Build & Test](https://github.com/DeltaLaboratory/go-imageserver/actions/workflows/go.yml/badge.svg)](https://github.com/DeltaLaboratory/go-imageserver/actions/workflows/go.yml)                | [![Docker Build & Push](https://github.com/DeltaLaboratory/go-imageserver/actions/workflows/docker.yml/badge.svg?branch=v0.1.0)](https://github.com/DeltaLaboratory/go-imageserver/actions/workflows/docker.yml)         |
| develop | [![CodeQL](https://github.com/DeltaLaboratory/go-imageserver/actions/workflows/codeql-analysis.yml/badge.svg?branch=develop)](https://github.com/DeltaLaboratory/go-imageserver/actions/workflows/codeql-analysis.yml) | [![Build & Test](https://github.com/DeltaLaboratory/go-imageserver/actions/workflows/go.yml/badge.svg?branch=develop)](https://github.com/DeltaLaboratory/go-imageserver/actions/workflows/go.yml) | [![Docker Build & Push](https://github.com/DeltaLaboratory/go-imageserver/actions/workflows/docker.yml/badge.svg?branch=v0.1.0-alpha.1)](https://github.com/DeltaLaboratory/go-imageserver/actions/workflows/docker.yml) |

## how to use
### run directly
First, install libaom and gcc as compile dependencies by `sudo apt update && sudo apt install -y libaom-dev build-essential`(debian-based linux) or `sudo yum update && sudo yum install -y libaom-devel gcc`(centos-based linux)\.
Next, clone this repo, and execute the program by `go run main.go`.
### run through docker
`docker pull ghcr.io/deltalaboratory/go-imageserver:0.1.0` and `docker run [some image dir or volume]:/go/src/app/images/ -p 8000:80 ghcr.io/deltalaboratory/go-imageserver:0.1.0`

## supported format
| upload | download |
|--------|----------|
| png    | webp     |
| jpeg   | avif     |
| webp   |
| tiff   |
| bmp    |
| gif    |