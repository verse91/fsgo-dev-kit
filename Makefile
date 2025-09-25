# Build the CLI tool
cli:
	go build -o fsgo .
	sudo cp fsgo /usr/local/bin/
	@echo "DONE"

# Build from cmd/fsgo (alternative entry point)
build:
	go build -o fsgo ./cmd/fsgo

# Install locally
install:
	go install .

# Clean build artifacts
clean:
	rm -f fsgo

# Test the project
test:
	go test ./...

# Run with go run
run:
	go run .

