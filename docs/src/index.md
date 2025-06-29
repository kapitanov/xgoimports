---
title: xgoimports
hide:
  - navigation
---
# xgoimports

A better goimports that keeps your imports sorted and grouped nicely

## Why?

As Go developers we often would like to have our imports grouped by origin:
standard library, then third-party libraries and at last - local packages.
But mainaining this order is a pain, especially when your imports are separated
by empty lines - and auto-importing in IDE is quite notorious to produce such a mess:

```go
// "github.com/myorg/myproject" is mean to be a local package.
import (
    "github.com/myorg/myproject/subpkg1"

    "github.com/myorg/myproject/subpkg2"
    "go.uber.org/atomic"

    "github.com/rs/zerolog/log"

    "fmt"
)
```

Here is where `xgoimports` comes in handy:

<div class="grid cards" markdown>

- With `go fmt` or `goimports`:

    ```go
    // "github.com/myorg/myproject" is mean to be
    // a local package.

    import (
        "github.com/myorg/myproject/subpkg1"

        "go.uber.org/atomic"

        "github.com/myorg/myproject/subpkg2"

        "github.com/rs/zerolog/log"

        "fmt"
    )
    ```

    - ☹️ Imports are not grouped by origin, it's a mess!
    - ☹️ Even `goimports -local github.com/myorg/myproject` doesn't help here.

- With `xgoimports`:

    ```go
    // "github.com/myorg/myproject" is mean to be
    // a local package.

    import (
        "fmt"

        "github.com/rs/zerolog/log"
        "go.uber.org/atomic"

        "github.com/myorg/myproject/subpkg1"
        "github.com/myorg/myproject/subpkg2"
    )
    ```

    - 😀 Imports are grouped by origin, so it's much easier to read and maintain.
    - 😀 Even if (_or when_) an IDE messes up with your imports - `xgoimports` will fix them at once!

</div>

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

=== "For your scripts"

    ```bash
    #!/bin/sh
    GO_FILES="$(find . -type f -name '*.go' -not -path "./vendor/*")"
    LOCAL_PACKAGES="github.com/yourusername/yourproject"

    goimports -local="$LOCAL_PACKAGES" -w -format-only $GO_FILES
    ```

=== "For your Makefile"

    ```makefile
    GO_FILES       := $(shell find . -type f -name '*.go' -not -path "./vendor/*")
    LOCAL_PACKAGES := github.com/yourusername/yourproject

    format:
        @xgoimports -local=$(LOCAL_PACKAGES) -w -format-only $(GO_FILES)
    ```

=== "For VSCode"

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

These import clauses will be left as-is:

```go
// This "import" clause won't be reformatted by xgoimports,
// however, the original behaviour of goimports is preserved.
import (
    "fmt"

    _ "github.com/myorg/myproject/subpkg2" // Has to be imported for side effects.

    "github.com/rs/zerolog/log"
    "go.uber.org/atomic"

    "github.com/myorg/myproject/subpkg1"
)
```

## Installation

We provide pre-built binaries for the most common platforms, so you don't need to build it from sources.
There are several ways to install `xgoimports`:

### Install via Homebrew (macOS)

You can install `xgoimports` via Homebrew on macOS. This is the recommended way to install it.

```bash
brew tap kapitanov/apps
brew install kapitanov/apps/xgoimports
```

### Install from Linix packages

We provide pre-built packages for the most common Linux distributions.
You can install `xgoimports` using the package manager of your distribution.

=== "DEB package (Ubuntu/Debian)"

    ```bash
    export VERSION="0.1.0" # replace with the actual version you want to install
    export ARCH="amd64"    # replace with the actual architecture (amd64, arm64, etc.)
    wget "https://github.com/kapitanov/xgoimports/releases/download/v${VERSION}/xgoimports_v${VERSION}_linux_${ARCH}.deb" \
        -O "xgoimports_v${VERSION}_linux_${ARCH}.deb"
    sudo dpkg -i "xgoimports_v${VERSION}_linux_${ARCH}.deb"
    ```

=== "RPM package (CentOS/RHEL/Fedora/AWS Linux)"

    ```bash
    export VERSION="0.1.0" # replace with the actual version you want to install
    export ARCH="amd64"    # replace with the actual architecture (amd64, arm64, etc.)
    wget "https://github.com/kapitanov/xgoimports/releases/download/v${VERSION}/xgoimports_v${VERSION}_linux_${ARCH}.rpm" \
        -O "xgoimports_v${VERSION}_linux_${ARCH}.rpm"
    sudo rpm -i "xgoimports_v${VERSION}_linux_${ARCH}.rpm"
    ```

=== "APK package (Alpine)"

    Run the following commands to install `xgoimports` on Alpine Linux:

    ```bash
    export VERSION="0.1.0" # replace with the actual version you want to install
    export ARCH="amd64"    # replace with the actual architecture (amd64, arm64, etc.)
    wget "https://github.com/kapitanov/xgoimports/releases/download/v${VERSION}/xgoimports_v${VERSION}_linux_${ARCH}.apk" \
        -O "xgoimports_v${VERSION}_linux_${ARCH}.apk"
    sudo apk add --allow-untrusted "xgoimports_v${VERSION}_linux_${ARCH}.apk"
    ```

### From pre-built binaries (for all platforms)

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

This project is licensed under the BSD-3-Clause License.

??? quote "License text"

    BSD 3-Clause License
    SPDX-License-Identifier: BSD-3-Clause

    Copyright 2009 The Go Authors.
    Copyright 2025 Albert Kapitanov.

    Redistribution and use in source and binary forms, with or without
    modification, are permitted provided that the following conditions are
    met:

    * Redistributions of source code must retain the above copyright
    notice, this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above
    copyright notice, this list of conditions and the following disclaimer
    in the documentation and/or other materials provided with the
    distribution.
    * Neither the name of Google LLC nor the names of its
    contributors may be used to endorse or promote products derived from
    this software without specific prior written permission.

    THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
    "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
    LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
    A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
    OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
    SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
    LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
    DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
    THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
    (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
    OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

!!! note "Note"

    This project is is a fork of the original `goimports` project,
    which was y developed by The Go Authors and is licensed under the BSD-3-Clause License as well.

    The original code is licensed under the BSD 3-Clause License,
    and portions of it have been modified for this project.

    No endorsement by Google or The Go Authors is implied.
