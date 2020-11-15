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
	@echo "run"
	@echo "test"

build: clean
	@echo "Building for ${GOOS}/${GOARCH}..."
	#@go build -o ${APP} *.go
	GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 \
	go build \
		-ldflags "-s -w -X ${PROJECT}/internal/version.Version=${VERSION} \
		-X ${PROJECT}/internal/version.Commit=${COMMIT} \
		-X ${PROJECT}/internal/version.BuildTime=${BUILD_TIME}" \
		-o bin/${APP} ${PROJECT}

container: build
	@sed                            \
	    -e 's|{GOOS}|$(GOOS)|g'     \
	    -e 's|{GOARCH}|$(GOARCH)|g' \
	    -e 's|{APP}|$(APP)|g'       \
	    -e 's|{PORT}|$(PORT)|g' 	\
	    Dockerfile.in > .dockerfile-$(APP)-$(GOOS)_$(GOARCH)
	@docker build -t ${APP}_v${VERSION} -f .dockerfile-$(APP)-$(GOOS)_$(GOARCH) .

run: container
	docker stop ${APP}_v${VERSION} || true && docker rm ${APP}_v${VERSION} || true
	docker run --name ${APP}_v${VERSION} -p ${PORT}:${PORT} \
		-e PORT=${PORT} -e APP=${APP} \
		-d ${APP}_v${VERSION}

fmt:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"gofmt -s -l {{.Dir}}"{{end}}' $(shell go list ${PROJECT}/...) | xargs -L 1 sh -c

vet:
	@echo "+ $@"
	@go vet $(shell go list ${PROJECT}/...)

test: clean fmt lint vet
	@echo "Run tests..."
	@go test --race ./...

cover:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"{{end}}' $(shell go list ${PROJECT}/...) | xargs -L 1 sh -c

clean:
	@echo "Cleaning"
	@go clean
	@rm -rf .dockerfile-*