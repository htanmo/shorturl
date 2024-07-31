# Default Recipe
default:
    @just --list

# Run all tests
test:
    @go test -v -cover ./...
