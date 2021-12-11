# go-imageserver
go-imageserver is an image server that automatically optimize images to webp and avif.

#how to use
first, install libaom and gcc for compile dependencies by `sudo apt update && sudo apt install -y libaom-dev build-essential`(debian-based linux) or `sudo yum update && sudo yum install -y libaom-devel gcc`(centos-based linux)\
and clone this repo, run `go run main.go` to start the server.