# go-imageserver
go-imageserver is an image server which automatically optimize non webp and avif images to webp and avif images.

## how to use
First, install libaom and gcc as compile dependencies by `sudo apt update && sudo apt install -y libaom-dev build-essential`(debian-based linux) or `sudo yum update && sudo yum install -y libaom-devel gcc`(centos-based linux)\.
Next, clone this repo, and execute the program by `go run main.go`.
