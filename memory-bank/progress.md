# Project Progress: md2docx

## Current Status (April 2025)

-   The core `md2docx` CLI tool is functional for basic Markdown-to-DOCX conversion using Pandoc.
-   It accepts input files and optional reference templates.
-   CI workflow (`ci.yml`) is mostly stable, successfully building and testing on Linux, macOS, and Windows.
-   Packaging scripts for Debian (`.deb`), macOS (`.tar.gz`), Windows (`.zip`), and NSIS installer exist within the `release.yml` workflow.
-   A Knowledge Base (`KB/`) exists documenting GitHub Actions concepts and project-specific workflows.
-   A Memory Bank (`memory-bank/`) has been initialized.

## What Works

-   Basic command-line execution: `md2docx input.md`.
-   Conversion using a reference template: `md2docx -template ref.docx input.md`.
-   Cross-platform builds via CI (`ci.yml`).
-   Windows shell integration via `.reg` file.
-   Installation via `go install` and building from source.

## Known Issues / What's Left

-   **Release Workflow (`release.yml`) is Broken:**
    -   **All recent release attempts have failed.**
    -   **Root Cause:** Incorrect version string extraction (`VERSION=${GITHUB_REF#refs/tags/}` results in empty string), leading to errors in packaging steps (e.g., `cp: cannot stat 'md2docx__linux_amd64'`).
    -   **Contributing Factors:** Uses outdated/archived GitHub Actions (`actions/create-release@v1`, `actions/upload-release-asset@v1`). See `KB/08-github-actions-version-audit.md`.
-   **Packaging Verification:** The packaging steps within `release.yml` haven't been successfully executed end-to-end due to the workflow failures. The generated packages (`.deb`, installer, etc.) need verification once the workflow is fixed.
-   **Homebrew Formula:** The Homebrew formula (`packaging/homebrew/md2docx.rb`) likely needs updates to align with fixed release assets/URLs once the release workflow is operational.
-   **Testing:** Unit and integration test coverage could be improved. Current tests might be minimal.

## Next Steps (Immediate Focus)

1.  **Fix `release.yml` Workflow:**
    -   Implement robust version handling using `${{ github.ref_name }}` and `workflow_dispatch` inputs.
    -   Update `actions/checkout` to `@v4`.
    -   Update `actions/setup-go` to `@v5`.
    -   Replace `actions/create-release@v1` and `actions/upload-release-asset@v1` with `softprops/action-gh-release@v1`.
    -   Standardize asset naming using the correct version variable.
2.  **Test the Fixed Release Workflow:** Trigger a release (e.g., by pushing a test tag) and verify successful execution and artifact generation.
3.  **Verify Release Assets:** Download and inspect the generated packages (`.deb`, `.zip`, `.tar.gz`, installer) to ensure they are correct.
