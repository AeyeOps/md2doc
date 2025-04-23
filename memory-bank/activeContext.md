# Active Context: md2docx (April 2025)

## Current Focus

The immediate priority is to **fix the broken GitHub Actions release workflow (`release.yml`)**. All recent attempts to create releases using this workflow have failed.

## Problem Diagnosis

Analysis of the workflow file (`release.yml`), execution logs (`gh run view 14605935691 --log`), and supporting documents (`codex.md`, `KB/08-github-actions-version-audit.md`) revealed the following key issues:

1.  **Faulty Version Extraction:** The mechanism `VERSION=${GITHUB_REF#refs/tags/}` is unreliable and produces an empty string, causing downstream filename errors (e.g., `md2docx__linux_amd64`) and packaging failures.
2.  **Outdated/Archived Actions:** The workflow uses:
    *   `actions/checkout@v3` (Latest: `v4`)
    *   `actions/setup-go@v4` (Latest: `v5`)
    *   `actions/create-release@v1` (Potentially outdated)
    *   `actions/upload-release-asset@v1` (**Archived/Unmaintained**)

## Action Plan

The agreed-upon plan to fix `release.yml` involves the following steps:

1.  **Refactor Version Handling:**
    *   Remove the shell-based `VERSION=${GITHUB_REF#refs/tags/}` line.
    *   Use `${{ github.ref_name }}` to get the version from tag pushes.
    *   Add a `workflow_dispatch` input named `version` for manual runs.
    *   Implement logic to select the correct version source based on the trigger event (`github.event_name`).
2.  **Update Actions:**
    *   Change `actions/checkout@v3` to `actions/checkout@v4`.
    *   Change `actions/setup-go@v4` to `actions/setup-go@v5`.
    *   Replace `actions/create-release@v1` and all `actions/upload-release-asset@v1` steps with `softprops/action-gh-release@v1`.
3.  **Standardize Asset Naming & Upload:**
    *   Ensure all build, packaging, and upload steps consistently use the correctly determined version variable.
    *   Configure `softprops/action-gh-release@v1` to upload all desired artifacts (`.deb`, `.zip`, `.tar.gz`, installer) using glob patterns.
4.  **Add Manual Trigger Input:**
    *   Define the `version` input within the `on.workflow_dispatch.inputs` section.

## Next Steps

-   Implement the planned changes to `.github/workflows/release.yml` using the `replace_in_file` tool.
-   Trigger the updated workflow (e.g., push a test tag) to verify the fix.
-   Inspect the generated release and artifacts for correctness.
