# Security Best Practices for GitHub Actions

## Understanding Security Risks

GitHub Actions workflows run with specific permissions and can access repository secrets. This powerful combination requires careful security consideration to prevent:

- **Secret leakage**: Accidental exposure of credentials or tokens
- **Supply chain attacks**: Compromised third-party actions
- **Privilege escalation**: Workflows gaining unwanted permissions
- **Resource abuse**: Actions consuming excessive resources
- **Injection attacks**: Command injection via user inputs

## Securing Workflow Secrets

GitHub provides a secure way to store and access sensitive information:

```yaml
jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - name: Deploy with secret
        env:
          API_TOKEN: ${{ secrets.API_TOKEN }}
        run: ./deploy.sh
```

Best practices for secrets management:

1. **Limit Scope**: Use environment-specific secrets
2. **Regular Rotation**: Change secret values periodically
3. **Least Privilege**: Grant minimal necessary permissions to secrets
4. **Audit Usage**: Regularly review where and how secrets are used

## Securing Third-Party Actions

When using actions from the marketplace or community:

```yaml
steps:
  # Secure usage with specific SHA
  - uses: actions/checkout@a12a394
  
  # Less secure - could change unexpectedly
  - uses: actions/checkout@v4
  
  # Least secure - could be completely different code
  - uses: actions/checkout@main
```

Security recommendations:

1. **Pin to SHA**: Use the full 40 character SHA hash when possible
2. **Verify Code**: Review the action's code before using it
3. **Trusted Authors**: Use actions from verified creators
4. **First-Party Actions**: Prefer official GitHub actions when available

## Limiting Workflow Permissions

By default, the `GITHUB_TOKEN` granted to workflows has access to repository contents and many API endpoints. Restrict this to the minimum required:

```yaml
permissions:
  contents: read
  issues: write
  # Other permissions set to none by default
```

For complete lockdown with opt-in permissions:

```yaml
permissions: {}  # No permissions by default

jobs:
  specific-job:
    permissions:
      contents: read
```

## Preventing Injection Attacks

User inputs must be treated carefully to prevent command injection:

```yaml
# Dangerous - direct command injection possible
- name: Run tool
  run: tool --flag ${{ github.event.inputs.parameter }}

# Better - uses GitHub's built-in context isolation
- name: Run tool safely
  run: |
    PARAM="${{ github.event.inputs.parameter }}"
    tool --flag "$PARAM"
```

Additional precautions:

1. **Input Validation**: Validate all inputs before use
2. **Avoid Direct Execution**: Never use user inputs directly in script execution
3. **Use Actions Toolkit**: Prefer the Actions toolkit functions which handle inputs securely

## OIDC Authentication for Cloud Providers

Use OpenID Connect (OIDC) to authenticate with cloud providers without storing long-lived credentials:

```yaml
permissions:
  id-token: write
  contents: read

jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
      - name: Configure AWS credentials
        uses: aws-actions/configure-aws-credentials@v2
        with:
          role-to-assume: arn:aws:iam::123456789012:role/my-github-actions-role
          aws-region: us-east-1
```

## Environment Protection Rules

Protect sensitive environments with required reviewers:

```yaml
jobs:
  deploy-production:
    runs-on: ubuntu-latest
    environment:
      name: production
      url: https://production.example.com
    steps:
      - name: Deploy
        run: ./scripts/deploy.sh production
```

Configure environment protection in repository settings:
1. Required reviewers before deployment
2. Wait timer to prevent accidental deployments
3. Branch restrictions to limit which branches can deploy

## Self-Hosted Runner Security

When using self-hosted runners:

1. **Isolation**: Use ephemeral runner containers or VMs
2. **Limited Access**: Only allow specific repositories to use runners
3. **Network Segmentation**: Isolate runners in dedicated network segments
4. **Regular Updates**: Keep runner software and host OS patched
5. **Minimal Tooling**: Install only necessary tools

## Dependency Supply Chain Security

Protect against compromised dependencies:

1. **Dependency Review**: Enable the dependency review action
2. **Dependabot Alerts**: Configure Dependabot to scan for vulnerabilities
3. **Dependency Verification**: Implement checksum validation

Example dependency review workflow:
```yaml
name: 'Dependency Review'
on: [pull_request]

permissions:
  contents: read

jobs:
  dependency-review:
    runs-on: ubuntu-latest
    steps:
      - name: 'Checkout Repository'
        uses: actions/checkout@v4
      - name: 'Dependency Review'
        uses: actions/dependency-review-action@v3
```

## Artifact Attestation

Establish provenance for build artifacts using SLSA (Supply chain Levels for Software Artifacts) framework:

```yaml
- name: Generate provenance
  uses: slsa-framework/slsa-github-generator@v1
  with:
    artifact-path: ./bin/app
    attestation-name: attestation.intoto.jsonl
```

## Best Practices Summary

1. **Principle of Least Privilege**: Grant minimal permissions to workflows and tokens
2. **Pin External Actions**: Use commit SHAs to reference third-party actions
3. **Secret Management**: Store secrets securely and limit their scope
4. **Input Validation**: Sanitize and validate all user inputs
5. **Environment Protection**: Implement approval workflows for sensitive environments
6. **Dependency Scanning**: Regularly scan dependencies for vulnerabilities
7. **Audit Logging**: Monitor workflow runs for suspicious activity
8. **Code Reviews**: Require reviews for workflow file changes
9. **Security Testing**: Include security tests in your CI/CD pipeline
10. **Documentation**: Maintain clear security documentation for your workflows