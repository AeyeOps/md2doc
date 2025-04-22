# Reusable Workflows and Shared Actions

## Creating Reusable Workflows

Reusable workflows allow you to define a workflow once and use it in multiple repositories, reducing duplication and maintenance overhead.

### Creating a Reusable Workflow

```yaml
# .github/workflows/reusable-go-build.yml
name: Reusable Go Build

on:
  workflow_call:
    inputs:
      go-version:
        required: false
        type: string
        default: '1.21'
      run-tests:
        required: false
        type: boolean
        default: true
    secrets:
      token:
        required: false

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      
      - name: Set up Go
        uses: actions/setup-go@v5
        with:
          go-version: ${{ inputs.go-version }}
          
      - name: Build
        run: go build -v ./...
        
      - name: Test
        if: ${{ inputs.run-tests }}
        run: go test -v ./...
```

### Using a Reusable Workflow

```yaml
# .github/workflows/main.yml
name: Main

on:
  push:
    branches: [ main ]

jobs:
  call-build-workflow:
    uses: ./.github/workflows/reusable-go-build.yml
    with:
      go-version: '1.21'
      run-tests: true
    secrets:
      token: ${{ secrets.GITHUB_TOKEN }}
```

## Creating Custom Actions

Custom actions help share common functionality across workflows.

### JavaScript Action

```yaml
# action.yml
name: 'Go Test Coverage'
description: 'Run tests and report coverage'
inputs:
  coverage-threshold:
    description: 'Minimum code coverage percentage required'
    required: false
    default: '80'
runs:
  using: 'node16'
  main: 'index.js'
```

```javascript
// index.js
const core = require('@actions/core');
const exec = require('@actions/exec');

async function run() {
  try {
    // Get inputs
    const threshold = core.getInput('coverage-threshold');
    
    // Run tests with coverage
    await exec.exec('go test -coverprofile=coverage.out ./...');
    
    // Process coverage data
    let output = '';
    await exec.exec('go tool cover -func=coverage.out', [], {
      listeners: {
        stdout: (data) => {
          output += data.toString();
        }
      }
    });
    
    // Extract total coverage
    const match = output.match(/total:\s+\(statements\)\s+(\d+\.\d+)%/);
    
    if (match && match[1]) {
      const coverage = parseFloat(match[1]);
      core.setOutput('coverage', coverage);
      
      if (coverage < parseFloat(threshold)) {
        core.setFailed(`Code coverage ${coverage}% is below the threshold ${threshold}%`);
      } else {
        core.info(`Code coverage: ${coverage}%`);
      }
    } else {
      core.setFailed('Could not determine code coverage');
    }
  } catch (error) {
    core.setFailed(error.message);
  }
}

run();
```

### Composite Action

```yaml
# action.yml
name: 'Go Setup and Cache'
description: 'Set up Go with dependency caching'
inputs:
  go-version:
    description: 'Go version to use'
    required: false
    default: '1.21'
runs:
  using: "composite"
  steps:
    - name: Set up Go
      uses: actions/setup-go@v5
      with:
        go-version: ${{ inputs.go-version }}
        cache: true
        
    - name: Verify dependencies
      run: go mod verify
      shell: bash
```

### Docker Container Action

```yaml
# action.yml
name: 'Go Cross-Compiler'
description: 'Cross-compile Go applications'
inputs:
  targets:
    description: 'Comma-separated list of OS/arch targets'
    required: true
  output-dir:
    description: 'Directory for compiled binaries'
    required: false
    default: 'dist'
runs:
  using: 'docker'
  image: 'Dockerfile'
  args:
    - ${{ inputs.targets }}
    - ${{ inputs.output-dir }}
```

```dockerfile
# Dockerfile
FROM golang:1.21

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
```

```bash
#!/bin/bash
# entrypoint.sh
set -e

TARGETS=$1
OUTPUT_DIR=$2

mkdir -p $OUTPUT_DIR

IFS=',' read -ra TARGET_ARRAY <<< "$TARGETS"
for target in "${TARGET_ARRAY[@]}"; do
  IFS='/' read -ra PARTS <<< "$target"
  OS=${PARTS[0]}
  ARCH=${PARTS[1]}
  
  echo "Building for $OS/$ARCH..."
  GOOS=$OS GOARCH=$ARCH go build -o "$OUTPUT_DIR/app_${OS}_${ARCH}" ./...
done
```

## Sharing Workflows and Actions Organization-Wide

### Repository Structure for Shared Actions

```
.github/
└── actions/
    ├── go-lint/
    │   ├── action.yml
    │   └── ... action files
    ├── go-build/
    │   ├── action.yml
    │   └── ... action files
    └── go-release/
        ├── action.yml
        └── ... action files
```

### Organization-Level Workflow Templates

Create `.github` repository in your organization with workflow templates:

```
.github/
└── workflow-templates/
    ├── go-ci.yml
    ├── go-ci.properties.json
    ├── go-release.yml
    └── go-release.properties.json
```

Example template definition:

```json
{
  "name": "Go CI Workflow",
  "description": "Go CI workflow template for building and testing applications.",
  "iconName": "go-icon",
  "categories": ["Go", "CI"],
  "filePatterns": ["go.mod$", "go.sum$"]
}
```

## Best Practices for Reusable Components

1. **Documentation**: Include comprehensive documentation for inputs, outputs, and usage examples
2. **Versioning**: Tag versions of your actions and reference specific tags in workflows
3. **Input Validation**: Validate inputs early in your action's execution
4. **Error Handling**: Provide clear error messages and handle failures gracefully
5. **Idempotency**: Ensure actions can be run multiple times without side effects
6. **Testing**: Create test workflows that use your actions in various scenarios
7. **Security**: Apply the principle of least privilege for any credentials used