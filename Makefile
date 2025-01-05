build:
	@go build -o bin/ecomgo cmd/main.go
test:
	@go test -v ./...

run: build
	@./bin/ecomgo