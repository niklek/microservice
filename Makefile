APP=service1
COMMIT=$(shell git rev-parse --short HEAD)

.PHONY: build
build: clean
	@echo "Building..."
	@go build -o ${APP} *.go

.PHONY: run
run:
	go run *.go

.PHONY: clean
clean:
	@echo "Cleaning"
	@go clean