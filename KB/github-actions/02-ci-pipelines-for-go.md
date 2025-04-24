# CI Pipelines for Go Applications

## Setting Up a Go CI Pipeline

GitHub Actions provides excellent support for building and testing Go applications. A comprehensive CI pipeline for Go projects typically includes the following stages:

1. **Checkout code**
2. **Setup Go environment**
3. **Dependency management**
4. **Code formatting verification**
5. **Static code analysis**
6. **Testing**
7. **Build**
8. **Artifact creation**

## Example Workflow

```yaml
name: Go CI

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  build:
    name: Build and Test
    runs-on: ubuntu-latest
    
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
          cache: true
      
      - name: Verify dependencies
        run: go mod verify
      
      - name: Install gofumpt
        run: go install mvdan.cc/gofumpt@latest
      
      - name: Check formatting
        run: |
          gofumpt -l .
          test -z "$(gofumpt -l .)"
      
      - name: Run vet
        run: go vet ./...
      
      - name: Install staticcheck
        run: go install honnef.co/go/tools/cmd/staticcheck@latest
      
      - name: Run staticcheck
        run: staticcheck ./...
      
      - name: Run tests
        run: go test -race -v ./...
      
      - name: Build
        run: go build -v ./...
```

## Multi-Version Testing with Matrix Strategy

Test your code against multiple Go versions to ensure compatibility:

```yaml
jobs:
  test:
    name: Test with Go ${{ matrix.go-version }}
    runs-on: ubuntu-latest
    strategy:
      matrix:
        go-version: ['1.19', '1.20', '1.21.x']
    
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go ${{ matrix.go-version }}
        uses: actions/setup-go@v5
        with:
          go-version: ${{ matrix.go-version }}
      
      - name: Run tests
        run: go test -v ./...
```

## Optimizing Go Builds with Caching

Improve build times by caching Go module dependencies:

```yaml
- name: Set up Go
  uses: actions/setup-go@v5
  with:
    go-version: '1.21'
    cache: true
    cache-dependency-path: go.sum
```

The `cache` parameter in `setup-go` creates and restores a cache of Go modules automatically.

## Code Coverage Reporting

Add coverage reporting to your workflow:

```yaml
- name: Run tests with coverage
  run: go test -race -coverprofile=coverage.txt -covermode=atomic ./...

- name: Upload coverage to Codecov
  uses: codecov/codecov-action@v3
  with:
    files: ./coverage.txt
```

## Custom Linting with golangci-lint

For comprehensive linting:

```yaml
- name: Run golangci-lint
  uses: golangci/golangci-lint-action@v3
  with:
    version: latest
    args: --timeout 3m
```

## Go Module Authentication

When working with private Go modules:

```yaml
- name: Configure Go module authentication
  run: git config --global url."https://${{ secrets.GH_ACCESS_TOKEN }}@github.com".insteadOf "https://github.com"
```

## Best Practices

1. **Test Coverage**: Aim for high test coverage and include integration tests
2. **Multiple Go Versions**: Test with the latest stable version and at least one previous version
3. **Linting**: Enforce code quality with static analysis tools
4. **Dependency Scanning**: Include vulnerability scanning for dependencies
5. **Documentation Generation**: Automatically generate and verify documentation
6. **Parallelism**: Use the matrix strategy to speed up testing across configurations