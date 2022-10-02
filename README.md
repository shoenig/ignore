ignore
======

Ignore unimportant error values

[![Go Reference](https://pkg.go.dev/badge/github.com/shoenig/ignore.svg)](https://pkg.go.dev/github.com/shoenig/ignore)
![GitHub](https://img.shields.io/github/license/shoenig/ignore.svg)

# Project Overview

Module `github.com/shoenig/ignore` provides a package with functions for ignoring
unimportant error values. These are useful for indicating errors that do not matter,
while also getting your IDE to stop highlighting ignored error returns.

# Getting Started

The `ignore` package can be installed by running
```bash
go get github.com/shoenig/ignore
```

#### Example Usage of ignore
```golang
// Close, the old way
// _ = closer.Close()

// Close, the new way
ignore.Close(closer)

// Drain, the old way
// _, _ = io.Copy(io.Discard, reader)
// _ = reader.Close()

// Drain, the new way
ignore.Drain(reader)
```

# Contributing

The `github.com/shoenig/ignore` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file an issue.

# License

The `github.com/shoenig/ignore` module is open source under the [BSD-3-Clause](LICENSE) license.
