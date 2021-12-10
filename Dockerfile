FROM golang:1.17-bullseye

WORKDIR /go/src/app
COPY . .

RUN apt -y update && apt install -y build-essential && apt install -y libaom-dev
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go-imageserver"]