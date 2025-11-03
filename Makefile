.PHONY: build install clean test run

# Build the application
build:
	go build -o show-port .

# Install the application
install:
	go install .

# Clean build artifacts
clean:
	rm -f show-port
	go clean

# Run tests
test:
	go test -v ./...

# Run the application
run:
	go run .

# Build for multiple platforms
build-all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o show-port-linux-amd64 .
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o show-port-linux-arm64 .
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o show-port-darwin-amd64 .
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o show-port-darwin-arm64 .
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o show-port-windows-amd64.exe .

# Format code
fmt:
	go fmt ./...

# Run linter
lint:
	golangci-lint run

# Show help
help:
	@echo "Available targets:"
	@echo "  build      - Build the application"
	@echo "  install    - Install the application"
	@echo "  clean      - Clean build artifacts"
	@echo "  test       - Run tests"
	@echo "  run        - Run the application"
	@echo "  build-all  - Build for all platforms"
	@echo "  fmt        - Format code"
	@echo "  lint       - Run linter"
	@echo "  help       - Show this help message"
