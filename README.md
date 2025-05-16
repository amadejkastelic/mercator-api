# Mercator API Client

The Mercator API Client is a Go library designed to interact with the [Mercator Online](https://www.mercatoronline.si/) API.

[![go report card](https://goreportcard.com/badge/github.com/amadejkastelic/mercator-api "go report card")](https://goreportcard.com/report/github.com/amadejkastelic/mercator-api)
[![CI status](https://github.com/amadejkastelic/mercator-api/actions/workflows/build.yaml/badge.svg?branch=main "test status")](https://github.com/amadejkastelic/mercator-api/actions)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/amadejkastelic/mercator-api?tab=doc)

## Overview

* Supports listing products and categories.
* Supports pagination and filtering.
* Provides a command-line interface (CLI).

## Running

To run the provided example, use the following command:
```bash
› nix run .#default -- --help
```

Find current price of Monster White Energy Drink:
```bash
› nix run .#default -- -query="Monster White"
Found 1 products
1. Energijski napitek, White, Monster, 0,5 l - 1.88€
```

## Development

```bash
› nix develop
```

## Build

```bash
› nix build
```

## Test

To run tests, use the following commands:
```bash
› nix develop --command go test ./...
```

To run all checks, including tests, use:
```bash
› nix flake check
```
