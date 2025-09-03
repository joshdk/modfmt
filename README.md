[![License][license-badge]][license-link]
[![Actions][github-actions-badge]][github-actions-link]
[![GoDoc][godoc-badge]][godoc-link]
[![Releases][github-release-badge]][github-release-link]

# modfmt

🗂️ Formatter for go.mod and go.work files

## Motivation

I was looking for a formatter for `go.mod` files and there didn't appear to be one that satisfied all requirements.

Consistent formatting is a given, but `modfmt` also makes sure to implement the following: 

- Supports formatting **both** `go.mod` and `go.work` files.
- Supports **all** the current `go.mod` and `go.work` directives.
  - See https://go.dev/ref/mod#go-mod-file.
  - See https://go.dev/ref/mod#go-work-file.
- Preserves both file header comments and directive comments.
- Can be used to verify that files are formatted inside a CI pipeline.

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
