SHELL := /bin/zsh

.PHONY: run test fmt lint

run:
	go run ./cmd/hello Youta

test:
	go test ./...

fmt:
	go fmt ./...

lint:
	go vet ./...
