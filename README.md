ignore
======

Ignore unimportant error values

[![Go Report Card](https://goreportcard.com/badge/go.gophers.dev/pkgs/ignore)](https://goreportcard.com/report/go.gophers.dev/pkgs/ignore)
[![Build Status](https://travis-ci.com/shoenig/ignore.svg?branch=master)](https://travis-ci.com/shoenig/ignore)
[![GoDoc](https://godoc.org/go.gophers.dev/pkgs/ignore?status.svg)](https://godoc.org/go.gophers.dev/pkgs/ignore)
![NetflixOSS Lifecycle](https://img.shields.io/osslifecycle/shoenig/ignore.svg)
![GitHub](https://img.shields.io/github/license/shoenig/ignore.svg)

# Project Overview

Module `go.gophers.dev/pkgs/ignore` provides a package with functions for ignoring
unimportant error values. These are useful for indicating errors that do not matter,
while also getting your IDE to stop highlighting ignored error returns.

# Getting Started

The `ignore` package can be installed by running
```bash
go get go.gophers.dev/pkgs/ignore
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

The `go.gophers.dev/pkgs/ignore` module is always improving with new features
and error corrections. For contributing bug fixes and new features please file an issue.

# License

The `go.gophers.dev/pkgs/ignore` module is open source under the [BSD-3-Clause](LICENSE) license.
