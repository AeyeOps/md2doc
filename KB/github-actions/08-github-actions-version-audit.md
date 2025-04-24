# GitHub Actions Version Audit (April 2025)

## Purpose

This document summarizes the status and latest versions of GitHub Actions used in the `md2doc` project workflows (`.github/workflows/ci.yml` and `.github/workflows/release.yml`) as of April 22, 2025. This information is intended to guide updates and ensure workflows use current, maintained actions.

## Actions Used and Status

### 1. `actions/checkout`

-   **Workflows:** `ci.yml`, `release.yml`
-   **Current Version Used:** `v3`
-   **Latest Major Version:** `v4`
-   **Status:** Maintained
-   **Recommendation:** Update to `@v4` for the latest features and security updates.

### 2. `actions/setup-go`

-   **Workflows:** `ci.yml`, `release.yml`
-   **Current Version Used:** `v4`
-   **Latest Major Version:** `v5`
-   **Status:** Maintained
-   **Recommendation:** Update to `@v5` for improved caching and compatibility.

### 3. `actions/create-release`

-   **Workflows:** `release.yml`
-   **Current Version Used:** `v1`
-   **Latest Major Version:** `v1` (Potentially outdated, see Status)
-   **Status:** Maintained, but often used in conjunction with the now-archived `upload-release-asset`. Modern alternatives combine creation and upload.
-   **Recommendation:** Replace with a combined action like `softprops/action-gh-release@v1` or `ncipollo/release-action@v1`.

### 4. `actions/upload-release-asset`

-   **Workflows:** `release.yml`
-   **Current Version Used:** `v1`
-   **Latest Major Version:** `v1` (Archived)
-   **Status:** **Archived and Unmaintained** (as of March 2021). The repository explicitly recommends using alternatives.
-   **Recommendation:** **Replace immediately**. Use a maintained action like `softprops/action-gh-release@v1` which handles both release creation and asset uploading.

## Summary of Recommendations

-   Update `actions/checkout` from `@v3` to `@v4`.
-   Update `actions/setup-go` from `@v4` to `@v5`.
-   Replace `actions/create-release@v1` and `actions/upload-release-asset@v1` in `release.yml` with a single, maintained action such as `softprops/action-gh-release@v1`. This simplifies the workflow and ensures use of supported tooling.
