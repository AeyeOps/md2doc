# md2docx

![CI](https://github.com/yourusername/md2docx/actions/workflows/ci.yml/badge.svg)
![Release](https://github.com/yourusername/md2docx/actions/workflows/release.yml/badge.svg)

 **md2docx** is a command-line tool to convert Markdown files to Microsoft Word `.docx` documents using [Pandoc](https://pandoc.org/). It also provides easy integration with the Windows Shell to add a right-click context menu option for generating Word documents.

 ## Features
 - Convert Markdown to `.docx` with a single command.
 - Optional reference `.docx` template for custom styles.
 - Windows Shell integration via a `.reg` file.

 ## Requirements
 - [Go](https://golang.org/dl/) (>= 1.20) for building the CLI.
 - [Pandoc](https://pandoc.org/installing.html) installed and available in your `PATH`.

 ## Installation

### 1. Using Go install
Ensure you have Go installed (>=1.20) and that your `$GOPATH/bin` (Linux/macOS) or `%USERPROFILE%\go\bin` (Windows) is in your `PATH`, then run:
```bash
go install github.com/yourusername/md2docx/cmd/md2docx@latest
```

### 2. Build from source
Clone the repository and build locally:
```bash
git clone https://github.com/yourusername/md2docx.git
cd md2docx
go build -o md2docx ./cmd/md2docx
```
For cross‑compiling to Windows on Linux/macOS:
```bash
GOOS=windows GOARCH=amd64 go build -o md2docx.exe ./cmd/md2docx
```
Copy the resulting binary (`md2docx` or `md2docx.exe`) to a directory in your `PATH`.

### 3. Pre‑built binaries
Download the latest release for your platform from the [GitHub Releases page](https://github.com/yourusername/md2docx/releases), extract it, and place the `md2docx` (or `md2docx.exe`) binary into a directory in your `PATH`.

 ## Usage

 Basic usage:
 ```bash
 md2docx [options] input.md
 ```

 Options:
 - `-template path` : Path to a reference `.docx` template file (to control styles, margins, and page size).
 - `-verbose`       : Enable verbose output.

 The output file will be created next to the input file with the same base name and a `.docx` extension.

 ## Windows Shell Integration

 Import the registry file to add a "Generate Word Document" context menu option for Markdown files:
 1. Double-click `windows-context-menu.reg` on Windows.
 2. Confirm the prompt to import into the registry.
 3. Right-click on a `.md` file and select **Generate Word Document**.

 See [`windows-context-menu.reg`](windows-context-menu.reg) for details.

## CI & Releases

We use GitHub Actions for continuous integration and automated releases.

- **CI**: Builds and tests on push/PR to `main` (Linux/macOS/Windows).
- **Releases**: Tag a new version to trigger a release workflow that cross-compiles binaries, packages them, and uploads to GitHub Releases.

To publish a new release:
```bash
git tag vX.Y.Z
git push origin vX.Y.Z
```

## Contributing
 Contributions are welcome! Please open issues and submit PRs on the GitHub repository.

 ## License
 This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
