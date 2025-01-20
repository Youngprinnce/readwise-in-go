run: build
	@./bin/readwise

build:
	@go build -o bin/readwise

test:
	@go test -v ./...
