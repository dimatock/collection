# Makefile for the collection project

.PHONY: all test bench lint tidy fmt

# Default target - runs the main quality assurance pipeline
all: tidy fmt lint test

# Format all go files
fmt:
	@echo "Formatting code..."
	@gofmt -s -w .

# Run all tests
test:
	@echo "Running tests..."
	@go test -v ./...

# Run all benchmarks
bench:
	@echo "Running benchmarks..."
	@go test -bench=. -benchmem ./...

# Run the linter
lint:
	@echo "Running linter..."
	@golangci-lint run

# Tidy up dependencies
tidy:
	@echo "Tidying go.mod..."
	@go mod tidy
