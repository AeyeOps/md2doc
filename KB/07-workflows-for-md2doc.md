# GitHub Actions Workflows for md2doc

This document provides tailored GitHub Actions workflow examples specifically for the md2doc project.

## CI Workflow

This workflow builds and tests the application on every push and pull request to the main branch:

```yaml
name: CI

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
      
      # Ensure pandoc is installed for integration tests
      - name: Install pandoc
        run: sudo apt-get update && sudo apt-get install -y pandoc
      
      - name: Verify dependencies
        run: go mod verify
      
      - name: Run tests
        run: go test -v ./...
      
      - name: Check formatting
        run: |
          if [ "$(gofmt -s -l . | wc -l)" -gt 0 ]; then
            gofmt -s -l .
            exit 1
          fi
      
      - name: Run vet
        run: go vet ./...
      
      - name: Build
        run: go build -v ./cmd/md2docx
```

## Cross-Platform Build Workflow

This workflow builds the application for multiple platforms:

```yaml
name: Cross-Platform Build

on:
  push:
    branches: [ main ]

jobs:
  build:
    name: Build ${{ matrix.os }}-${{ matrix.arch }}
    runs-on: ubuntu-latest
    strategy:
      matrix:
        os: [linux, darwin, windows]
        arch: [amd64, arm64]
        exclude:
          # Skip Windows ARM64 build
          - os: windows
            arch: arm64
    
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Build
        env:
          GOOS: ${{ matrix.os }}
          GOARCH: ${{ matrix.arch }}
        run: |
          # Add .exe suffix for Windows binaries
          if [ "${{ matrix.os }}" = "windows" ]; then
            go build -o "md2docx_${{ matrix.os }}_${{ matrix.arch }}.exe" ./cmd/md2docx
          else
            go build -o "md2docx_${{ matrix.os }}_${{ matrix.arch }}" ./cmd/md2docx
          fi
      
      - name: Upload artifacts
        uses: actions/upload-artifact@v3
        with:
          name: md2docx-${{ matrix.os }}-${{ matrix.arch }}
          path: md2docx_${{ matrix.os }}_${{ matrix.arch }}*
          retention-days: 5
```

## Release Workflow

This workflow creates a new release when a tag is pushed:

```yaml
name: Release

on:
  push:
    tags:
      - 'v*'

jobs:
  goreleaser:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4
        with:
          fetch-depth: 0
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v4
        with:
          distribution: goreleaser
          version: latest
          args: release --clean
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

Required `.goreleaser.yml` configuration file:

```yaml
# .goreleaser.yml
builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - windows
      - darwin
    goarch:
      - amd64
      - arm64
    ignore:
      - goos: windows
        goarch: arm64
    main: ./cmd/md2docx
    ldflags:
      - -s -w -X main.version={{.Version}}

archives:
  - format_overrides:
      - goos: windows
        format: zip
    files:
      - README.md
      - LICENSE
      - windows-context-menu.reg

checksum:
  name_template: 'checksums.txt'

snapshot:
  name_template: "{{ incpatch .Version }}-next"

changelog:
  sort: asc
  filters:
    exclude:
      - '^docs:'
      - '^test:'
```

## Docker Image Workflow

This workflow builds and publishes a Docker image:

```yaml
name: Docker

on:
  push:
    branches: [ main ]
    tags: [ 'v*' ]

jobs:
  docker:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout
        uses: actions/checkout@v4
      
      - name: Docker meta
        id: meta
        uses: docker/metadata-action@v5
        with:
          images: ghcr.io/aeyeops/md2doc
          tags: |
            type=ref,event=branch
            type=ref,event=pr
            type=semver,pattern={{version}}
            type=semver,pattern={{major}}.{{minor}}
      
      - name: Set up Docker Buildx
        uses: docker/setup-buildx-action@v3
      
      - name: Login to GitHub Container Registry
        uses: docker/login-action@v3
        with:
          registry: ghcr.io
          username: ${{ github.repository_owner }}
          password: ${{ secrets.GITHUB_TOKEN }}
      
      - name: Build and push
        uses: docker/build-push-action@v5
        with:
          context: .
          push: ${{ github.event_name != 'pull_request' }}
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}
          cache-from: type=gha
          cache-to: type=gha,mode=max
```

Required `Dockerfile`:

```dockerfile
# Dockerfile
FROM golang:1.21 AS builder
WORKDIR /app
COPY . .
RUN go build -o md2docx ./cmd/md2docx

FROM ubuntu:22.04
RUN apt-get update && apt-get install -y pandoc ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app/md2docx /usr/local/bin/md2docx
ENTRYPOINT ["md2docx"]
```

## Documentation Build Workflow

This workflow builds and deploys documentation to GitHub Pages:

```yaml
name: Documentation

on:
  push:
    branches: [ main ]
    paths:
      - 'docs/**'
      - 'README.md'
      - '.github/workflows/docs.yml'

jobs:
  deploy:
    runs-on: ubuntu-latest
    permissions:
      contents: write
    steps:
      - uses: actions/checkout@v4
      
      - name: Setup Python
        uses: actions/setup-python@v4
        with:
          python-version: '3.x'
      
      - name: Install dependencies
        run: |
          python -m pip install --upgrade pip
          pip install mkdocs mkdocs-material
      
      - name: Build documentation
        run: mkdocs build
      
      - name: Deploy
        uses: peaceiris/actions-gh-pages@v3
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          publish_dir: ./site
```

## Integration Test Workflow

This workflow runs integration tests using real markdown files:

```yaml
name: Integration Tests

on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]

jobs:
  integration:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: '1.21'
      
      - name: Install pandoc
        run: sudo apt-get update && sudo apt-get install -y pandoc
      
      - name: Build
        run: go build -o md2docx ./cmd/md2docx
      
      - name: Run integration tests
        run: |
          ./md2docx examples/sample.md
          if [ ! -f examples/sample.docx ]; then
            echo "Failed to generate DOCX file"
            exit 1
          fi
          
          # Test with template
          ./md2docx -template examples/template.docx examples/sample.md
          if [ ! -f examples/sample.docx ]; then
            echo "Failed to generate DOCX file with template"
            exit 1
          fi
```

## Dependabot Configuration

Create `.github/dependabot.yml` to automatically update dependencies:

```yaml
version: 2
updates:
  - package-ecosystem: "gomod"
    directory: "/"
    schedule:
      interval: "weekly"
    open-pull-requests-limit: 10
    
  - package-ecosystem: "github-actions"
    directory: "/"
    schedule:
      interval: "weekly"
    open-pull-requests-limit: 10
```

## Complete CI/CD Pipeline

A comprehensive approach combining the workflows above into a complete pipeline:

1. **CI** runs on every push and pull request
2. **Cross-Platform Build** creates artifacts for all supported platforms
3. **Integration Tests** verify functionality with real documents
4. **Release** creates official releases with versioned binaries
5. **Docker** publishes container images for each version
6. **Documentation** keeps the GitHub Pages site updated
7. **Dependabot** ensures dependencies stay current

This covers everything needed for a professional, automated development workflow for the md2doc project.