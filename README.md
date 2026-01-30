# Updater

This is a tool used to fetch and update project data.

## Usage

**Requirement:**

- `go` >= `1.25.5`

1. Download `updater` latest:

```bash
go install github.com/universtar-org/updater@latest
```

> The `latest` could be replaced other version that can be found in [Tag Page](https://github.com/universtar-org/updater/tags), like `v1.0.0`.

2. Execute the command:

```bash
$(go env GOPATH)/bin/updater /path/to/project/data
```

## Acknowledgement

- [`go-yaml`](https://github.com/goccy/go-yaml): YAML support for the Go language.
