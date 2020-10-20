export GO111MODULE=on

.PHONY: build
build:
	go build -o bin/myserver cmd/myserver/main.go

.PHONY: generate-pb
generate-pb:
	go generate app/pb/gen.go

.PHONY: deps
deps:
	go mod tidy