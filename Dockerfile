FROM golang:1.17-bullseye AS build

WORKDIR /go/src/app
COPY . .

RUN apt -y update && apt install -y build-essential libaom-dev
RUN go get -d -v ./... \
    && go build -o application ./...

FROM debian:bullseye-slim
WORKDIR /
COPY statics ./statics
COPY --from=build /go/src/app/application .
RUN apt update -y \
    && apt install -y libaom-dev
RUN mkdir -p images \
    && mkdir -p config \
    && chmod +x ./application
ENTRYPOINT ["/application"]