build:
	@go build -o bin/quickrelay cmd/main.go

run:
	@go run cmd/main.go

clean:
	@rm -rf bin

test:
	@go test -v -failfast ./...
.PHONY: run clean test
