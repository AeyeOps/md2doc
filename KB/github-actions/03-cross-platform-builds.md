# Cross-Platform Builds & Releases for Go

## Cross-Compiling Go Applications

Go's design simplifies cross-compilation for multiple operating systems and architectures. GitHub Actions can leverage this capability to build binaries for various targets in parallel.

## Matrix Strategy for Cross-Platform Builds

The following example builds a Go application for multiple OS and architecture combinations:

```yaml
name: Cross-Platform Build

on:
  push:
    tags:
      - 'v*'

jobs:
  build:
    name: Build ${{ matrix.os }}-${{ matrix.arch }}
    runs-on: ubuntu-latest
    strategy:
      matrix:
        os: [linux, darwin, windows]
        arch: [amd64, arm64]
        exclude:
          # Exclude combinations that aren't relevant
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
```

## Creating GitHub Releases with Built Binaries

Automatically create GitHub releases when a new tag is pushed:

```yaml
jobs:
  # ... previous build job
  
  release:
    name: Create Release
    runs-on: ubuntu-latest
    needs: build
    
    steps:
      - name: Download all artifacts
        uses: actions/download-artifact@v3
        
      - name: Create Release
        id: create_release
        uses: softprops/action-gh-release@v1
        with:
          name: Release ${{ github.ref_name }}
          draft: false
          prerelease: false
          generate_release_notes: true
          files: |
            md2docx-linux-amd64/md2docx_linux_amd64
            md2docx-linux-arm64/md2docx_linux_arm64
            md2docx-darwin-amd64/md2docx_darwin_amd64
            md2docx-darwin-arm64/md2docx_darwin_arm64
            md2docx-windows-amd64/md2docx_windows_amd64.exe
```

## Packaging Binaries for Distribution

Package binaries with relevant files (README, LICENSE) for better user experience:

```yaml
- name: Package Binaries
  run: |
    mkdir -p release
    for dir in md2docx-*; do
      os_arch=${dir#md2docx-}
      cp $dir/md2docx_* release/
      cd release
      
      # Create zip archives for Windows, tar.gz for others
      if [[ "$os_arch" == "windows-"* ]]; then
        zip "md2docx_${os_arch}.zip" "md2docx_${os_arch}.exe" ../README.md ../LICENSE
      else
        tar -czf "md2docx_${os_arch}.tar.gz" "md2docx_${os_arch}" ../README.md ../LICENSE
      fi
      
      cd ..
    done
    
- name: Upload packaged binaries
  uses: softprops/action-gh-release@v1
  with:
    files: release/*
```

## Building with CGO Enabled

For cross-compilation with CGO dependencies, use Docker containers with the necessary toolchains:

```yaml
- name: Set up cross-compilation for CGO
  uses: docker://docker.elastic.co/beats-dev/golang-crossbuild:1.21.0
  with:
    args: >
      --build-cmd "go build -o md2docx_linux_arm64 ./cmd/md2docx" 
      -p "linux/arm64"
```

Alternatively, use a marketplace action designed for this purpose:

```yaml
- name: CGO Cross Compilation
  uses: crazy-max/ghaction-xgo@v3
  with:
    xgo_version: latest
    go_version: 1.21
    dest: build
    prefix: md2docx
    targets: windows/amd64,linux/amd64,linux/arm64,darwin/amd64,darwin/arm64
    v: true
    x: false
    race: false
    ldflags: -s -w
    buildmode: default
```

## SHA256 Checksums for Verification

Add checksums to your releases for security:

```yaml
- name: Generate SHA256 checksums
  run: |
    cd release
    sha256sum * > checksums.txt
    
- name: Upload checksums
  uses: softprops/action-gh-release@v1
  with:
    files: release/checksums.txt
```

## Integrating with goreleaser

[GoReleaser](https://goreleaser.com/) automates the entire release process:

```yaml
- name: Run GoReleaser
  uses: goreleaser/goreleaser-action@v4
  with:
    version: latest
    args: release --clean
  env:
    GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

With a `.goreleaser.yml` configuration file:

```yaml
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

## Best Practices for Cross-Platform Builds

1. **Testing on All Platforms**: Run tests across all target platforms
2. **Versioning**: Use semantic versioning for releases
3. **Build Metadata**: Include version information in binaries using ldflags
4. **Artifact Naming**: Use consistent naming conventions (app_os_arch)
5. **Documentation**: Include platform-specific instructions in README
6. **Signing**: Sign your binaries for platforms that support it (Windows, macOS)
7. **Verification**: Include checksums for all distributed binaries