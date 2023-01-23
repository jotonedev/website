.PHONY: build style test clean tool lint help

all: build style

build:
	@go build -v .

style:
	sass style/default.scss static/style.css

test:
	go test -v ./...

tool:
	go vet ./...; true
	gofmt -w .

lint:
	golint ./...

clean:
	rm -rf go-gin-example
	go clean -i .

help:
	@echo "make: compile packages and dependencies"
	@echo "make build: compile packages and dependencies"
	@echo "make test: run tests"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"