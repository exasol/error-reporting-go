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

The optional third parameter for `parameter(placeholder, value, description)` will be used to generate a parameter description.

The builder automatically quotes parameters (depending on the type of the parameter) with single quotes.

### Mitigations

The mitigations describe those actions the user can follow to overcome the error, and are specified as follows:

```go
renderedString := error_reporting_go.ExaError("E-TEST-2").Message("Too little disk space.").Mitigation("Delete something.").String()
```

Result: `E-TEST-2: Too little disk space. Delete something.`

## Additional Resources

* [Changelog](doc/changes/changelog.md)