# UniverStar Tools

This is a tool set used in `UniverStar` project.

## Requirement

- `go` >= `1.25.5`

## Usage

### Updater

Updater is used to update project information, such as update time, number of stars, etc.

1. Download latest version

   ```bash
   curl -L -o updater https://github.com/universtar-org/tools/releases/latest/download/updater-linux-amd64
   chmod +x updater
   ```

2. Update project data with the path to data files

   ```bash
   ./updater ./data/projects
   ```

### Checker

Checker is used to check whether a repository is accessible.

1. Download latest version

   ```bash
   curl -L -o checker https://github.com/universtar-org/tools/releases/latest/download/checker-linux-amd64
   chmod +x checker
   ```

2. Check data files under a given directory

   ```bash
   ./checker ./data/projects
   ```

### Unique

Unique is used to check duplicated username across all repository within `universtar-org`.

1. Download latest version

   ```bash
   curl -L -o unique https://github.com/universtar-org/tools/releases/latest/download/unique-linux-amd64
   chmod +x unique
   ```

2. Check the uniqueness of a user

   ```bash
   ./unique [username]
   ```

## Acknowledgement

- [`go-yaml`](https://github.com/goccy/go-yaml): YAML support for the Go language.
- [`tint`](https://github.com/lmittmann/tint): Log beautifier.
