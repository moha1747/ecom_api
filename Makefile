# Define the binary output path
BINARY=bin/ecom_api

# Default target
all: build

# Build the project
build:
	@go build -o $(BINARY) cmd/main.go

# Run tests
test:
	@go test -v ./...

# Run the compiled binary
run: build
	@./$(BINARY)

# Clean the build output
clean:
	@rm -f $(BINARY)
	@echo "Cleaned up build artifacts"

.PHONY: all build test run clean
