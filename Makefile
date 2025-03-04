# Project Name
APP_NAME := stockexchange

# Output Directory
BIN_DIR := bin

# Go Build Flags
GOFLAGS := -ldflags="-s -w"

# Run tests
test:
	@go test -v ./...

# Lint code
lint:
	@golangci-lint run

# Build for the current OS/Arch
build:
	@echo "Building $(APP_NAME) for current OS..."
	@go build $(GOFLAGS) -o $(BIN_DIR)/$(APP_NAME) .

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	@rm -rf $(BIN_DIR)

# Install dependencies
deps:
	@go mod tidy
