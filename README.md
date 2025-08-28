# safebox

A lightweight, secure toolkit for managing secrets and sensitive files with a simple and predictable workflow.

## Introduction

Safebox helps teams store, retrieve, and share secrets safely across environments. It is designed to be:
- Simple: minimal commands and clear outputs
- Secure: sensible defaults, encryption at rest and in transit
- Portable: works well in local development, CI/CD, and containerized environments

Common use cases:
- Managing API keys, tokens, and certificates
- Encrypting configuration files for deployments
- Securely sharing secrets among team members

## Installation

Choose one of the following methods.

- Prebuilt binaries (recommended):
    1. Download the latest release for your OS from the Releases page.
    2. Add the binary to your PATH (e.g., /usr/local/bin).

- Build from source:
    - Requirements: Go 1.24 or newer
    - Steps:
        1. Clone the repository
        2. Run: `go build -o safebox ./cmd/safebox`
        3. Move `safebox` to a directory in your PATH

- Docker:
    - Pull: `docker pull <your-registry>/safebox:<version>`
    - Run: `docker run --rm -it -v "$PWD":/work -w /work <your-registry>/safebox:<version> safebox --help`

## Usage

Below are common commands to get started. Replace placeholders (like <name>) with your values.

- Initialize a workspace/vault:
  ```
  safebox init --path .safebox
  ```

- Add or update a secret:
  ```
  safebox put --key <NAME> --value <SECRET_VALUE>
  ```
  Or from a file:
  ```
  safebox put --key <NAME> --file ./path/to/secret.txt
  ```

- Retrieve a secret:
  ```
  safebox get --key <NAME>
  ```
  Output to file:
  ```
  safebox get --key <NAME> --out ./secret.txt
  ```

- List secrets (metadata only):
  ```
  safebox list
  ```

- Remove a secret:
  ```
  safebox delete --key <NAME>
  ```

- Encrypt/decrypt a file (ad‑hoc):
  ```
  safebox encrypt --in ./config.yaml --out ./config.yaml.enc
  safebox decrypt --in ./config.yaml.enc --out ./config.yaml
  ```

- Show help:
  ```
  safebox --help
  safebox <command> --help
  ```

Tips:
- Use environment variables for non-interactive usage in CI (e.g., SAFebox_* variables).
- Commit only encrypted artifacts to version control; never commit plaintext secrets.