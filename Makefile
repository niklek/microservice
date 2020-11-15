PROJECT?=github.com/niklek/microservice
APP?=service1
GOOS?=linux
GOARCH?=amd64
PORT?=8000
VERSION?=0.0.1
COMMIT?=$(shell git rev-parse --short HEAD)
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')

.PHONY: build container run test clean

help:
	@echo "Makefile commands:"
	@echo "build"
	@echo "run"
	@echo "test"

build: clean
	@echo "Building..."
	#@go build -o ${APP} *.go
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 \
	go build \
		-ldflags "-s -w -X ${PROJECT}/internal/version.Version=${VERSION} \
		-X ${PROJECT}/internal/version.Commit=${COMMIT} \
		-X ${PROJECT}/internal/version.BuildTime=${BUILD_TIME}" \
		-o bin/${APP} ${PROJECT}/*.go

container: build
	docker build --build-arg PORT=${PORT} --build-arg APP=${APP} -t ${APP}:${VERSION} .

run: container
	docker run --name ${APP}:${VERSION} -p ${PORT}:${PORT} \
		-e "PORT=${PORT}" \
		-d ${APP}:${VERSION}

test:
	@echo "Run tests..."
	@go test --race ./...

clean:
	@echo "Cleaning"
	@go clean