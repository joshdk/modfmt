[![License][license-badge]][license-link]
[![Actions][github-actions-badge]][github-actions-link]
[![GoDoc][godoc-badge]][godoc-link]
[![Releases][github-release-badge]][github-release-link]

# modfmt

🗂️ Formatter for go.mod and go.work files

## What is this?

This repository provides `modfmt`, a super simple formatter for `go.mod` and `go.work` files.

Additionally, it implements some specific features:

- Consistent ordering of sections and directives.
- Supports formatting of **both** `go.mod` and `go.work` files.
- Supports **all** of the current `go.mod` and `go.work` directives.
  - See https://go.dev/ref/mod#go-mod-file.
  - See https://go.dev/ref/mod#go-work-file.
- Preserves both file header comments and directive comments.
- Can be used in a CI pipeline to verify that files are formatted.
- Can be used as a library with minimal dependencies.
- Can be used as an `analysis.Analyzer`. (planned)

## Formatting

Each of the various `go.mod` and `go.work` directives are combined into a unified block, consistently sorted, and rendered along with any associated comments. The ordering of directive blocks was based off of ecosystem conventions.

### Ordering of `go.mod` directives

| Section              | Explanation                                                                                              |
|----------------------|----------------------------------------------------------------------------------------------------------|
| `// Header comments` | All header or unattached comments.                                                                       |
| `module …`           | The [module](https://go.dev/ref/mod#go-mod-file-module) directive.                                       |
| `go …`               | The [go](https://go.dev/ref/mod#go-mod-file-go) directive.                                               |
| `toolchain …`        | The [toolchain](https://go.dev/ref/mod#go-mod-file-toolchain) directive.                                 |
| `godebug (…)`        | A block of [godebug](https://go.dev/ref/mod#go-mod-file-godebug) directives.                             |
| `retract (…)`        | A block of [retract](https://go.dev/ref/mod#go-mod-file-retract) directives.                             |
| `require (…)`        | A block of [require](https://go.dev/ref/mod#go-mod-file-require) directives.                             |
| `require (…)`        | A block of [require](https://go.dev/ref/mod#go-mod-file-require) directives. (for indirect dependencies) |
| `ignore (…)`         | A block of [ignore](https://go.dev/ref/mod#go-mod-file-ignore) directives.                               |
| `exclude (…)`        | A block of [exclude](https://go.dev/ref/mod#go-mod-file-exclude) directives.                             |
| `replace (…)`        | A block of [replace](https://go.dev/ref/mod#go-mod-file-replace) directives.                             |
| `replace (…)`        | A block of [replace](https://go.dev/ref/mod#go-mod-file-replace) directives. (for local replacements)    |
| `tool (…)`           | A block of [tool](https://go.dev/ref/mod#go-mod-file-tool) directives.                                   |

### Ordering of `go.work` directives

| Section              | Explanation                                                                                            |
|----------------------|--------------------------------------------------------------------------------------------------------|
| `// Header comments` | All header or unattached comments.                                                                     |
| `go …`               | The [go](https://go.dev/ref/mod#go-work-file-go) directive.                                            |
| `toolchain …`        | The [toolchain](https://go.dev/ref/mod#go-work-file-toolchain) directive.                              |
| `godebug (…)`        | A block of [godebug](https://go.dev/ref/mod#go-work-file-godebug) directives.                          |
| `use (…)`            | A block of [use](https://go.dev/ref/mod#go-work-file-use) directives.                                  |
| `replace (…)`        | A block of [replace](https://go.dev/ref/mod#go-work-file-replace) directives.                          |
| `replace (…)`        | A block of [replace](https://go.dev/ref/mod#go-work-file-replace) directives. (for local replacements) |

## Installation

### Release artifact

Binaries for various architectures are published on the [releases][github-release-link] page.

The latest release can be installed by running:

```shell
OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m | sed 's/x86_64/amd64/; s/aarch64/arm64/')
wget -O modfmt.tar.gz https://github.com/joshdk/modfmt/releases/latest/download/modfmt-${OS}-${ARCH}.tar.gz
tar -xf modfmt.tar.gz
sudo install modfmt /usr/local/bin/modfmt
```

### Brew

Release binaries are also available via [Brew](https://brew.sh).

The latest release can be installed by running:

```shell
brew tap joshdk/tap
brew install joshdk/tap/modfmt
```

### Go install

Installation can also be done directly from this repository.

The latest commit can be installed by running:

```shell
go install github.com/joshdk/modfmt@master
```

## Usage

### Showing unformatted files

Show unformatted `go.mod` or `go.work` files in the current directory:

```shell
modfmt
```

Show unformatted files anywhere under the directory `pkg`:

```shell
modfmt pkg/...
```

List unformatted filenames anywhere under the directory `pkg`:

```shell
modfmt -l pkg/...
```

### Fixing unformatted files

Format and update all files under the current directory:

```shell
modfmt -w ./...
```

> [!IMPORTANT]  
> You should always run `go mod tidy` prior to `modfmt`.

Exit with an error if any files were unformatted.

### Verifying that files are formatted

```shell
modfmt -c ./...
```

> [!TIP]
> This command should be run in CI during a linting pass.

## License

This code is distributed under the [MIT License][license-link], see [LICENSE.txt][license-file] for more information.

---

<p align="center">
  Created by <a href="https://github.com/joshdk">Josh Komoroske</a> ☕
</p>

[github-actions-badge]:  https://github.com/joshdk/modfmt/actions/workflows/build.yaml/badge.svg
[github-actions-link]:   https://github.com/joshdk/modfmt/actions/workflows/build.yaml
[github-release-badge]:  https://img.shields.io/github/release/joshdk/modfmt/all.svg
[github-release-link]:   https://github.com/joshdk/modfmt/releases
[godoc-badge]:           https://pkg.go.dev/badge/github.com/joshdk/modfmt/pkg/modfmt
[godoc-link]:            https://pkg.go.dev//github.com/joshdk/modfmt/pkg/modfmt
[license-badge]:         https://img.shields.io/badge/license-MIT-green.svg
[license-file]:          https://github.com/joshdk/modfmt/blob/master/LICENSE.txt
[license-link]:          https://opensource.org/licenses/MIT
