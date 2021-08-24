# Go Error Reporting

This project contains a Go-Builder for Exasol error messages.

## Usage

### Simple Messages

```go
renderedString := error_reporting_go.ExaError("E-TEST-1").Message("Something went wrong.").String()
```

Result: `E-TEST-1: Something went wrong.`

### Parameters

You can specify placeholders in the message and fill them up with parameters values:

```go
renderedString := error_reporting_go.ExaError("E-TEST-2").Message("Unknown input {{input}}.").Parameter("input", "unknown").String()
```

Result: `E-TEST-2: Unknown input 'unknown'.`

`ParameterWithDescription(name, value, description)` method also gets a parameter description which will be used to generate a parameter description for the error-catalog.

The builder automatically quotes parameters with single quotes.
If you want to avoid quotes, use the `|uq` suffix in the correspondent placeholder:

```go
renderedString := error_reporting_go.ExaError("E-TEST-2").Message("Unknown input {{input|uq}}.").Parameter("input", 2).String()
```
result: `E-TEST-2: Unknown input 2.`

### Mitigations  

The mitigations describe those actions the user can follow to overcome the error, and are specified as follows:

```go
renderedString := error_reporting_go.ExaError("E-TEST-3").Message("Too little disk space.").Mitigation("Delete something.").String()
```

Result: `E-TEST-3: Too little disk space. Delete something.`

## Additional Resources

* [Changelog](doc/changes/changelog.md)