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

# Run the compiled binary with arguments


migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@migrate go run cmd/migrate/main.go up

migrate-down:
	@migrate go run cmd/migrate/main.go down

# Clean the build output
clean:
	@rm -f $(BINARY)
	@echo "Cleaned up build artifacts"

.PHONY: all build test run clean
