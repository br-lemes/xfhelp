.PHONY: build release test version

build: test
	@go build -ldflags "-s -w"

release: version build
	@go run ./tools/release/main.go

test:
	@go test ./...

version: test
	@go run ./tools/version/main.go
