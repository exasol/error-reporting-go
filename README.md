# Go Error Reporting

This project contains a Go-Builder for Exasol error messages.

[![Go Reference](https://pkg.go.dev/badge/github.com/exasol/error-reporting-go.svg)](https://pkg.go.dev/github.com/exasol/error-reporting-go)
[![Build](https://github.com/exasol/error-reporting-go/actions/workflows/build.yml/badge.svg)](https://github.com/exasol/error-reporting-go/actions/workflows/build.yml)

## Usage

### Including Go Error Reporting

Add this to your `go.mod`:

```
require (
    github.com/exasol/error-reporting-go v0.1.1
)
```

Then import the library in your `.go` files:

```go
import (
    exaerror "github.com/exasol/error-reporting-go"
)
```

### Simple Messages

```go
renderedString := exaerror.New("E-TEST-1").Message("Something went wrong.").String()
```

Result: `E-TEST-1: Something went wrong.`

### As native go error

```go
err := exaerror.New("E-TEST-1").Message("Something went wrong.")
fmt.Println(err)  // fmt package can print errors automatically 
fmt.Println(err.Error())  // Print error message explicit via Error() function
```

Result: `E-TEST-1: Something went wrong.`

### Parameters

You can specify placeholders in the message and fill them up with parameters values:

```go
renderedString := exaerror.New("E-TEST-2").Message("Unknown input {{input}}.").Parameter("input", "unknown").String()
```

or inline

```go
renderedString := exaerror.New("E-TEST-2").Messagef("Unknown input {{input}}.", "unknown").String()
```

Result: `E-TEST-2: Unknown input 'unknown'.`

`ParameterWithDescription(name, value, description)` method also gets a parameter description which will be used to generate a parameter description for the error-catalog.

The builder automatically quotes parameters with single quotes.
If you want to avoid quotes, use the `|uq` suffix in the correspondent placeholder:

```go
renderedString := exaerror.New("E-TEST-2").Message("Unknown input {{input|uq}}.").Parameter("input", 2).String()
```
result: `E-TEST-2: Unknown input 2.`

### Mitigations  

The mitigations describe those actions the user can follow to overcome the error, and are specified as follows:

```go
renderedString := exaerror.New("E-TEST-3").Message("Too little disk space.").Mitigation("Delete something.").String()
```

Result: `E-TEST-3: Too little disk space. Delete something.`

## Additional Resources

* [Changelog](doc/changes/changelog.md)