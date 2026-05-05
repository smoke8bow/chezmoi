# Contributing to chezmoi

Thank you for your interest in contributing to chezmoi! This document provides
guidelines and information to help you contribute effectively.

## Code of Conduct

Please read and follow our [Code of Conduct](.github/CODE_OF_CONDUCT.md).

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.21 or later
- [Git](https://git-scm.com/)
- [golangci-lint](https://golangci-lint.run/) for linting

### Setting Up Your Development Environment

1. Fork the repository on GitHub.
2. Clone your fork locally:
   ```sh
   git clone https://github.com/<your-username>/chezmoi.git
   cd chezmoi
   ```
3. Add the upstream remote:
   ```sh
   git remote add upstream https://github.com/twpayne/chezmoi.git
   ```
4. Install dependencies:
   ```sh
   go mod download
   ```

## Making Changes

### Branching

Create a new branch for your changes:
```sh
git checkout -b feat/my-new-feature
```

Use a descriptive branch name that reflects the change you are making.

### Coding Standards

- Follow standard Go conventions and idioms.
- Run `golangci-lint run` before submitting and fix any reported issues.
- Ensure all tests pass with `go test ./...`.
- Write tests for new functionality.
- Keep commits focused and atomic.

### Commit Messages

We follow [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>(<scope>): <short summary>

[optional body]

[optional footer]
```

Common types: `feat`, `fix`, `docs`, `refactor`, `test`, `chore`.

Examples:
- `feat(cmd): add --dry-run flag to apply command`
- `fix(template): handle missing variable gracefully`
- `docs: update installation instructions`

## Submitting a Pull Request

1. Push your branch to your fork:
   ```sh
   git push origin feat/my-new-feature
   ```
2. Open a Pull Request against the `main` branch of this repository.
3. Fill in the PR template with a clear description of your changes.
4. Ensure CI checks pass.
5. Respond to review feedback promptly.

## Reporting Issues

When filing a bug report, please include:
- chezmoi version (`chezmoi --version`)
- Operating system and version
- Steps to reproduce the issue
- Expected vs. actual behavior
- Relevant configuration or log output (redact sensitive data)

## Questions

For general questions, please open a [GitHub Discussion](https://github.com/twpayne/chezmoi/discussions)
rather than an issue.
