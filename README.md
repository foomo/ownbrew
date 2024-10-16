# Ownbrew

[![Build Status](https://github.com/foomo/ownbrew/actions/workflows/test.yml/badge.svg?branch=main&event=push)](https://github.com/foomo/ownbrew/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/foomo/ownbrew)](https://goreportcard.com/report/github.com/foomo/ownbrew)
[![godoc](https://godoc.org/github.com/foomo/ownbrew?status.svg)](https://godoc.org/github.com/foomo/ownbrew)
[![goreleaser](https://github.com/foomo/ownbrew/actions/workflows/release.yml/badge.svg)](https://github.com/foomo/ownbrew/actions)

> Your local project package manager

## Installing

Install the latest release of the cli:

````bash
$ brew update
$ brew install foomo/tap/ownbrew
````

## Usage

Add a `ownbrew.yaml` configuration

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/foomo/ownbrew/v0.1.0/ownbrew.yaml
version: '1.0'

binDir: "bin"
tapDir: ".ownbrew/tap"
tempDir: ".ownbrew/tmp"
cellarDir: ".ownbrew/bin"
packages:
  ## https://github.com/golangci/golangci-lint/releases
  - name: golangci-lint
    tap: foomo/tap/golangci/golangci-lint
    version: 1.61.0
  ## https://github.com/go-courier/husky/releases
  - name: husky
    tap: foomo/tap/go-courier/husky
    version: 1.8.1
```

Install your local packages

```shell
$ ownbrew install
```

Add the `bin` folder to your `$PATH`

```shell
$ export PATH=bin:$PATH
```

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.
