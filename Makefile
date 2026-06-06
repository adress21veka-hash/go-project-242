.PHONY: lint lint-fix

lint:
	golangci-lint run ./...

build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size