# xgoimports

![GitHub branch check runs](https://img.shields.io/github/check-runs/kapitanov/xgoimports/master?label=build)
![GitHub Release](https://img.shields.io/github/v/release/kapitanov/xgoimports)
![GitHub License](https://img.shields.io/github/license/kapitanov/xgoimports)

A better goimports that keeps your imports sorted and grouped nicely

## Why?

As Go developers we often would like to have our imports grouped by origin:

- standard library
- third-party libraries
- local packages

Built-in `go fmt` command doesn't support this formatting style, but thankfully, `goimports` does!
Still, it doesn't force-group imports, so you can end up with a mess like this:

```go
import (
    "fmt"

    "github.com/some/package"

    "os"

    "github.com/another/package"

    "github.com/some/package/subpackage"

    "github.com/third/package"

    "github.com/third/package/subpackage"
)
```

Some IDEs (like GoLand) might produce this mess while trying to generate imports for you on the fly.

This is were `xgoimports` comes in handy. It will sort and group your imports automatically, so you can end up with something like this:

```go
// Assuming that
//   - "github.com/some/package" is a local package,
//   - "github.com/another/package" is a third-party package,
//   - "github.com/third/package" is a third-party package as well,
//   - and the rest are standard library packages.
import (
    "fmt"
    "os"

    "github.com/another/package"
    "github.com/third/package"
    "github.com/third/package/subpackage"

    "github.com/some/package"
    "github.com/some/package/subpackage"
)
```

## Usage

`xgoimports` is a drop-in replacement for `goimports`, so you can use it in the same way:

```bash
xgoimports -w yourfile.go
```

It's recommmended to use it with `-local` flag to specify your local packages, so they are grouped together:

```bash
xgoimports -local=github.com/yourusername/yourproject -w -format-only yourfile.go
```

Often, you would want to run it on all your Go files in the project:

```bash
xgoimports -local=github.com/yourusername/yourproject -w -format-only $(find . -type f -name '*.go' -not -path "./vendor/*")
```

### Handy snippets

- For your scripts:

  ```bash
  #!/bin/sh
  GO_FILES="$(find . -type f -name '*.go' -not -path "./vendor/*")"
  LOCAL_PACKAGES="github.com/yourusername/yourproject"

  goimports -local="$LOCAL_PACKAGES" -w -format-only $GO_FILES
  ```

- For your Makefile:

  ```makefile
  GO_FILES       := $(shell find . -type f -name '*.go' -not -path "./vendor/*")
  LOCAL_PACKAGES := github.com/yourusername/yourproject

  format:
  	@xgoimports -local=$(LOCAL_PACKAGES) -w -format-only $(GO_FILES)
  ```

- For VSCode:

  ```json
  {
    "editor.formatOnSave": true,
    "go.formatTool": "custom",
    "go.alternateTools": {
      "customFormatter": "xgoimports"
    },
    "go.formatFlags": [
      "-local=\"github.com/yourusername/yourproject\"",
      "-w",
      "-format-only"
    ]
  }
  ```

### Current limitations

`xgoimports` is unable to properly format imports if there are any comments in the import block.
These import clauses will be left as-is.

## Installation

We provide pre-built binaries for the most common platforms, so you don't need to build it from sources.
There are several ways to install `xgoimports`:

### Install via Homebrew (macOS)

You can install `xgoimports` via Homebrew on macOS. This is the recommended way to install it.

```bash
brew tap kapitanov/apps
brew install kapitanov/apps/xgoimports
```

### Install from deb package (Ubuntu/Debian)

```bash
export VERSION="0.1.0" # replace with the actual version you want to install
export ARCH="amd64"    # replace with the actual architecture (amd64, arm64, etc.)
wget "https://github.com/kapitanov/xgoimports/releases/download/v${VERSION}/xgoimports_v${VERSION}_linux_${ARCH}.deb" \
    -O "xgoimports_v${VERSION}_linux_${ARCH}.deb"
sudo dpkg -i "xgoimports_v${VERSION}_linux_${ARCH}.deb"
```

### Install from rpm package (CentOS/RHEL/Fedora/AWS Linux)

```bash
export VERSION="0.1.0" # replace with the actual version you want to install
export ARCH="amd64"    # replace with the actual architecture (amd64, arm64, etc.)
wget "https://github.com/kapitanov/xgoimports/releases/download/v${VERSION}/xgoimports_v${VERSION}_linux_${ARCH}.rpm" \
    -O "xgoimports_v${VERSION}_linux_${ARCH}.rpm"
sudo rpm -i "xgoimports_v${VERSION}_linux_${ARCH}.rpm"
```

### Install from apk package (Alpine)

Run the following commands to install `xgoimports` on Alpine Linux:

```bash
export VERSION="0.1.0" # replace with the actual version you want to install
export ARCH="amd64"    # replace with the actual architecture (amd64, arm64, etc.)
wget "https://github.com/kapitanov/xgoimports/releases/download/v${VERSION}/xgoimports_v${VERSION}_linux_${ARCH}.apk" \
    -O "xgoimports_v${VERSION}_linux_${ARCH}.apk"
sudo apk add --allow-untrusted "xgoimports_v${VERSION}_linux_${ARCH}.apk"
```

### From Releases (for all platforms)

You can download the latest release of `xgoimports` from the [Releases](https://github.com/kapitanov/xgoimports/releases) page.
Just pick the version you want and the appropriate binary for your operating system and architecture.

### From sources (for all platforms)

You can install `xgoimports` via `go get`:

```bash
go install github.com/kapitanov/xgoimports@latest
```

Note that you need to have Go installed on your system to use this method - and you must have the `$GOPATH/bin` in your `PATH`.
Since you most definitely are a Go developer - you are likely to have all these prerequisites already.

## License

This project is licensed under the BSD-3-Clause License - see the [LICENSE](LICENSE) file for details.

Note that this project is a fork of the original `goimports` project,
which is licensed under the BSD-3-Clause License as well.
Please refer to the [NOTICE](NOTICE) file for more details.

## Acknowledgments

This project is a derivative work based on source code from the Go programming language toolchain,
originally developed by The Go Authors.
The original code is licensed under the BSD 3-Clause License,
and portions of it have been modified for this project.
No endorsement by Google or The Go Authors is implied.
