# Active Context: md2docx (April 2025)

## Current Focus

The task to **populate the Knowledge Base (KB) with Azure CLI reference documentation has been stopped** by the user. The next focus should be addressing the issues with the GitHub Actions `release.yml` workflow.

## Azure CLI KB Task Details

-   **Goal:** Fetch Azure CLI reference documentation and store it in the project's Knowledge Base (`KB/`).
-   **Source:** Microsoft Learn, starting from the A-Z index (`https://learn.microsoft.com/en-us/cli/azure/reference-index?view=azure-cli-latest`).
-   **Method:** Use the `firecrawl_scrape` tool to fetch content as Markdown, focusing on the main content area of each command group's page.
-   **Structure:** Create a dedicated subdirectory `KB/azure-cli/` and store one Markdown file per top-level command group (e.g., `KB/azure-cli/vm.md`, `KB/azure-cli/storage.md`).
-   **Status (as of stopping):**
    -   Directory `KB/azure-cli/` created.
    -   The following command group files were fetched and saved before stopping:
        -   `KB/azure-cli/acat.md`
        -   `KB/azure-cli/account.md`
        -   `KB/azure-cli/acr.md`
        -   `KB/azure-cli/ad.md`
        -   `KB/azure-cli/advisor.md`
        -   `KB/azure-cli/afd.md`
        -   `KB/azure-cli/ai-examples.md`
        -   `KB/azure-cli/aks.md`
        -   `KB/azure-cli/aksarc.md`
        -   `KB/azure-cli/akshybrid.md`
        -   `KB/azure-cli/alerts-management.md`
        -   `KB/azure-cli/alias.md`
        -   `KB/azure-cli/amlfs.md`
        -   `KB/azure-cli/ams.md`
        -   `KB/azure-cli/aosm.md`
    -   The process was stopped after saving `KB/azure-cli/aosm.md`.

## Next Steps

-   **Fix `release.yml` Workflow:** Address the previously identified issues with the GitHub Actions `release.yml` workflow (faulty version extraction, outdated actions). This is now the primary focus.
