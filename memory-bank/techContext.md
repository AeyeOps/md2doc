# Technical Context: md2docx

## Core Technologies

-   **Programming Language:** Go (Version >= 1.20 required for building from source, as specified in README).
-   **Conversion Engine:** Pandoc (External dependency, must be installed and in PATH).
-   **CI/CD Platform:** GitHub Actions.
-   **Version Control:** Git / GitHub.

## Development Environment

-   **Go Toolchain:** Standard Go compiler and tools (`go build`, `go test`, `go mod`, `gofmt`).
-   **Text Editor/IDE:** Any standard editor suitable for Go development (e.g., VS Code with Go extension).
-   **Operating System:** Development can occur on Linux, macOS, or Windows. CI tests run on all three.

## Dependencies

-   **Runtime:**
    -   `pandoc`: The core external dependency required for the tool to function. Its availability and version can impact conversion results.
-   **Build-time (Go Modules):**
    -   Currently, the `cmd/md2docx/main.go` uses only standard Go library packages (`flag`, `fmt`, `os`, `os/exec`, `path/filepath`). No external Go modules are listed in `go.mod` beyond the module definition itself.
-   **CI/CD Environment:**
    -   GitHub Actions runners (Ubuntu, Windows, macOS).
    -   Specific GitHub Actions (e.g., `actions/checkout`, `actions/setup-go`).
    -   Packaging tools installed within the workflow (`dpkg-dev`, `zip`, `tar`, `nsis`).

## Technical Constraints & Considerations

-   **Pandoc Dependency:** The tool's functionality is entirely dependent on a correctly installed and accessible `pandoc` executable. Errors in Pandoc installation or execution will cause `md2docx` to fail.
-   **Pandoc Version:** While not explicitly enforced, different Pandoc versions might produce slightly different output or support different Markdown features. Consistency relies on the user's Pandoc installation.
-   **GitHub Actions Runner Environment:** Packaging steps (Debian, NSIS) rely on tools (`dpkg-dev`, `nsis`) being installable on the `ubuntu-latest` runner used in the release workflow.
-   **Release Workflow Fragility:** The current release workflow (`release.yml`) has known issues (version parsing, outdated actions) making it unreliable (See `KB/08-github-actions-version-audit.md` and `codex.md`).

## Tooling Usage Patterns

-   `go build`: Used for compiling the application.
-   `go test`: Used for running unit/integration tests (though current tests might be minimal).
-   `gofmt`: Used for code formatting checks in CI.
-   `go vet`: Used for static analysis checks in CI.
-   `gh` (GitHub CLI): Used manually (or potentially in future workflows) for interacting with GitHub Actions runs.
-   `dpkg-dev`, `tar`, `zip`, `makensis`: Used within the release workflow for packaging.
