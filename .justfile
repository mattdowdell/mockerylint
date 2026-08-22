# https://just.systems/man/en/

[private]
default:
    @just --list

# Run all recipes.
all: tidy vendor fmt lint unit

# Tidy all dependencies.
[group('dependencies')]
tidy: tidy-root tidy-testdata

# Tidy the root module dependencies.
[group('dependencies')]
tidy-root:
    go mod tidy

# Tidy the testdata module dependencies.
[group('dependencies')]
tidy-testdata:
    cd testdata && go mod tidy

# Vendor dependencies.
[group('dependencies')]
vendor:
    go mod vendor

# Run all formatters.
[group('formatters')]
fmt: fmt-go fmt-just fmt-yaml

# Run the Go formatter.
[group('formatters')]
fmt-go:
    golangci-lint fmt

# Run the Justfile formatter.
[group('formatters')]
fmt-just:
    just --unstable --fmt

# Run the YAML formatter.
[group('formatters')]
fmt-yaml:
    yamlfmt .

# Run all linters.
[group('linters')]
lint: lint-go

# Run the Go linter.
[group('linters')]
lint-go:
    golangci-lint run

# Run the unit tests.
[group('tests')]
unit:
    go test -count=1 -cover ./...
