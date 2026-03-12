# xfhelp

A CLI tool for helping with XFCE configuration management.

## Overview

xfhelp interacts with `xfconf-query` to manage XFCE settings in a
structured and reproducible way. It uses a built-in schema to define
which channels and properties are supported, along with their expected
types, serving as the single source of truth for validation.

## Requirements

- `xfconf-query` available in `PATH`

## Installation

Build and install directly:

```sh
go install github.com/br-lemes/xfhelp@latest
```

Or clone the source code first:

```sh
git clone https://github.com/br-lemes/xfhelp
cd xfhelp
go build
```

## Commands

### export

Reads current XFCE settings and outputs them as JSON. By default, only
properties covered by the schema are included.

```sh
xfhelp export
xfhelp export --untracked  # properties not covered by the schema instead
```

### import

Applies settings from a JSON file to the current XFCE configuration.
Validates all channels, properties and types against the schema before
making any changes.

```sh
xfhelp import settings.json
xfhelp import --dry-run settings.json  # show pending changes without applying
```

### outputs

Lists available XFCE display outputs (monitors, screens). Shows both
default options ("Automatic", "Primary") and currently active display
devices detected by XFCE.

```sh
xfhelp outputs
```

### schema

Prints the JSON Schema of all supported channels and properties. Useful
for validating a settings file externally or understanding what xfhelp
tracks.

```sh
xfhelp schema
```

## Development

```sh
make test     # run tests
make build    # build binary (runs tests first)
make version  # compute and update version from git log
make release  # tag and publish a new release (runs version and build first)
```

Version numbers are derived automatically from the git commit history
following [Conventional Commits](https://www.conventionalcommits.org/):
breaking changes bump the major version, `feat` bumps minor, and `fix`
bumps patch.

## Contributing

Contributions are welcome! Feel free to open issues or pull requests.

## License

BSD Zero Clause License. See [LICENSE](LICENSE) file.
