# AGENTS.md

## Build Commands

- `go build` - Build the project
- `go install` - Install the binary
- `go test ./...` - Run all tests in the project
- `go test -v ./internal/io` - Run tests for io package with verbose output
- `go test ./internal/utils -v` - Run tests for utils package with verbose output
- `go test -run TestReadYaml` - Run a single test function

## Lint Commands

- `gofmt -d .` - Check formatting with gofmt
- `go vet` - Run go vet for static analysis
- `golint ./...` - Run golint (if available)

## Code Style Guidelines

### Imports

- Standard library imports first
- Then external dependencies
- Then local package imports
- Import groups separated by blank lines

### Formatting

- Use `gofmt` for formatting
- Follow Go's standard formatting conventions
- No trailing whitespace
- Use tabs for indentation

### Naming Conventions

- Use camelCase for variables and functions
- Use PascalCase for exported functions and types
- Use descriptive names; avoid abbreviations
- Use `TODO` comments for incomplete implementations

### Error Handling

- Handle errors explicitly with `if err != nil`
- Return errors from functions instead of panicking
- Use `errors.New()` for creating new errors
- Use `fmt.Errorf()` for formatted error messages

### Types

- Use explicit types rather than `var`
- Use `const` for constants
- Use `type` for type definitions

### Testing

- Use table-driven tests with named test cases
- Use `t.Run()` for sub-tests
- Use `cmp.Diff()` for test result comparison
- All tests must pass for code to be accepted

### General Guidelines

- Keep functions short and focused
- Use clear, descriptive comments
- Use godoc for package and exported function documentation
- Avoid global variables when possible
- Use context for timeouts and cancellation

