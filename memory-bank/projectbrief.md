# Project Brief: md2docx

## Core Goal

To provide a simple, fast, and reliable command-line tool (`md2docx`) for converting Markdown files (`.md`) into Microsoft Word documents (`.docx`).

## Key Objectives

-   **Conversion:** Leverage Pandoc for accurate Markdown-to-DOCX conversion.
-   **Simplicity:** Offer a straightforward CLI interface with minimal configuration.
-   **Speed:** Ensure fast conversion times suitable for quick documentation tasks.
-   **Cross-Platform:** Provide usable binaries for Windows, macOS, and Linux.
-   **Integration:** Offer optional Windows shell integration for ease of use.
-   **Customization:** Allow users to specify reference `.docx` templates for styling.

## Scope

-   Develop a Go-based CLI application.
-   Implement core conversion logic using `os/exec` to call Pandoc.
-   Handle command-line arguments for input file, template (optional), and verbosity.
-   Set up CI/CD pipelines using GitHub Actions for building, testing, and releasing.
-   Provide packaging solutions (Debian, Homebrew, Windows Installer, archives).
-   Maintain basic project documentation (README, CHANGELOG).
-   Establish a Knowledge Base (KB) for GitHub Actions implementation details.
