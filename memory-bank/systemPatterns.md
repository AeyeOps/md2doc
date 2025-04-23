# System Patterns: md2docx

## Core Architecture

`md2docx` follows a simple **CLI Wrapper Pattern** around the external `pandoc` executable.

```mermaid
graph LR
    User --> CLI[md2docx CLI (Go)]
    CLI -- Executes --> Pandoc(Pandoc Process)
    CLI -- Input File --> Pandoc
    CLI -- Optional Template --> Pandoc
    Pandoc -- Output File --> Filesystem
```

## Key Technical Decisions

-   **Language Choice (Go):** Chosen for its cross-compilation capabilities, performance, and suitability for building simple, self-contained CLI tools.
-   **Conversion Engine (Pandoc):** Leveraged as the underlying engine for Markdown-to-DOCX conversion due to its robustness, extensive format support, and maturity. `md2docx` acts as a user-friendly interface to Pandoc for this specific task.
-   **Dependency Management (Go Modules):** Standard Go modules (`go.mod`) are used for managing any potential Go dependencies (though the current implementation has minimal external Go dependencies).
-   **Automation (GitHub Actions):** Used for CI (building, testing across platforms) and CD (creating releases, packaging binaries).
-   **Packaging:** Multiple formats are provided (archives, `.deb`, NSIS installer, Homebrew formula) to cater to different user installation preferences across platforms.

## Implementation Patterns

-   **External Process Execution:** The Go `os/exec` package is used to invoke the `pandoc` command with appropriate arguments derived from CLI flags and input files.
-   **Command-Line Argument Parsing:** Go's standard `flag` package is used to parse user-provided options (`-template`, `-verbose`) and the input filename.
-   **File System Interaction:** Standard Go packages (`os`, `path/filepath`) are used for checking file existence, determining output paths, etc.
-   **Cross-Compilation:** Go's built-in cross-compilation features (`GOOS`, `GOARCH` environment variables) are utilized within the GitHub Actions release workflow to generate binaries for Linux, macOS, and Windows.

## Workflow Patterns (GitHub Actions)

-   **CI Workflow (`ci.yml`):** Matrix strategy used to build and test on `ubuntu-latest`, `windows-latest`, `macos-latest`.
-   **Release Workflow (`release.yml`):**
    -   Triggered by version tags (`v*.*.*`) and manual dispatch (`workflow_dispatch`).
    -   Builds binaries for multiple targets (`linux_amd64`, `windows_amd64`, `darwin_amd64`).
    -   Packages artifacts (`.deb`, `.zip`, `.tar.gz`, `.exe` installer).
    -   Uses (currently outdated/archived) actions for creating GitHub releases and uploading assets. **(Identified for refactoring)**.
    -   Relies on shell scripting for version extraction and packaging steps. **(Identified for refactoring)**.
