# Makefile for golf shot data processor

# Go parameters
GOCMD=go
GOTEST=$(GOCMD) test
GOVET=$(GOCMD) vet
GOMOD=$(GOCMD) mod
GORUN=$(GOCMD) run
GOFMT=$(GOCMD) fmt
GOLINT=staticcheck

# Main package path
MAIN_PACKAGE=./main.go

# Find all directories with go files
GO_DIRS=$(shell find . -type f -name '*.go' -not -path "./vendor/*" -print0 | xargs -0 -n1 dirname | sort -u)

# Targets
.PHONY: all test vet fmt lint clean run tidy help

all: test vet fmt lint

# Run tests
test:
	@echo "Running tests..."
	@for dir in $(GO_DIRS); do \
		echo "Testing $$dir..."; \
		(cd $$dir && $(GOTEST) -v); \
	done

# Run go vet
vet:
	@echo "Running go vet..."
	@for dir in $(GO_DIRS); do \
		echo "Vetting $$dir..."; \
		(cd $$dir && $(GOVET)); \
	done

# Format code
fmt:
	@echo "Formatting code..."
	@for dir in $(GO_DIRS); do \
		echo "Formatting $$dir..."; \
		(cd $$dir && $(GOFMT)); \
	done

# Run golint
lint:
	@echo "Running golint..."
	@for dir in $(GO_DIRS); do \
		echo "Linting $$dir..."; \
		(cd $$dir && $(GOLINT)); \
	done

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	@find . -type f -name '*.test' -delete

# Run the application
run:
	@echo "Running the application..."
	$(GORUN) $(MAIN_PACKAGE)

# Tidy and verify dependencies
tidy:
	@echo "Tidying and verifying module dependencies..."
	$(GOMOD) tidy
	$(GOMOD) verify

# Help target
help:
	@echo "Available targets:"
	@echo "  test    - Run tests in all packages"
	@echo "  vet     - Run go vet on all packages"
	@echo "  fmt     - Format all Go files"
	@echo "  lint    - Run golint on all packages"
	@echo "  clean   - Remove build artifacts"
	@echo "  run     - Run the application"
	@echo "  tidy    - Tidy and verify module dependencies"
	@echo "  all     - Run tests, vet, fmt, and lint"
	@echo "  help    - Show this help message"