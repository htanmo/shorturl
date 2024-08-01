# Default Recipe
default:
    @just --list

# Run all tests
test:
    @go test -v -cover ./...

# Build the project
build:
    @go build -o bin/shorturl

# Run the project
run: build
    @./bin/shorturl

# Clean the project
clean:
    @rm -rf bin/