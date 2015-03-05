# Migrator
[![GoDoc](https://godoc.org/github.com/c4milo/migrator?status.svg)](https://godoc.org/github.com/c4milo/migrator)
[![Build Status](https://travis-ci.org/c4milo/migrator.svg?branch=master)](https://travis-ci.org/c4milo/migrator)

Opinionated database migration library for Go applications.

### Supported databases
* Postgres


When building your project using this library, make sure  you pass build tags to compile only the driver you want to use. Example: `go build -tags postgres` or `go test -tags postgres`