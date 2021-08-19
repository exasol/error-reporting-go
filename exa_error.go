package error_reporting_go

import (
	"fmt"
	"reflect"
	"strings"
)

type ErrorMessageBuilder struct {
	errorCode  string
	message    string
	parameters [][]string
	mitigation string
}

func ExaError(errorCode string) *ErrorMessageBuilder {
	return &ErrorMessageBuilder{
		errorCode: errorCode,
	}
}

func (builder *ErrorMessageBuilder) Message(message string) *ErrorMessageBuilder {
	builder.message = message
	return builder
}

func (builder *ErrorMessageBuilder) Parameter(arguments ...interface{}) *ErrorMessageBuilder {
	validateArguments(arguments)
	builder.parameters = append([][]string{{fmt.Sprintf("%v", arguments[0]), fmt.Sprintf("%v", arguments[1])}})
	return builder
}

func (builder *ErrorMessageBuilder) Mitigation(mitigation string) *ErrorMessageBuilder {
	builder.mitigation = mitigation
	return builder
}

func validateArguments(arguments []interface{}) {
	if len(arguments) < 2 || len(arguments) > 3 {
		panic("Parameter function accepts 2 or 3 arguments: parameter name, parameter value, (optional) description")
	}
	if getInterfaceType(arguments[0]) != "string" {
		panic("Parameter function's first argument 'parameter name' must be a string")
	}
	if len(arguments) == 3 && getInterfaceType(arguments[2]) != "string" {
		panic("Parameter function's third argument 'description' must be a string")
	}
}

func getInterfaceType(argument interface{}) string {
	typeOf := reflect.TypeOf(argument)
	if typeOf == nil {
		return "nil"
	} else {
		return typeOf.String()
	}
}

func (builder *ErrorMessageBuilder) String() string {
	var stringBuilder strings.Builder
	formattedMessage := formatMessage(builder)
	stringBuilder.WriteString(fmt.Sprintf("%s: %s", builder.errorCode, formattedMessage))
	if len(builder.mitigation) > 0 {
		stringBuilder.WriteString(fmt.Sprintf(" %s", builder.mitigation))
	}
	return stringBuilder.String()
}

func formatMessage(builder *ErrorMessageBuilder) string {
	var formattedMessage = builder.message
	for _, parameter := range builder.parameters {
		formattedMessage = strings.Replace(formattedMessage, "{{"+parameter[0]+"}}", "'"+parameter[1]+"'", -1)
	}
	return formattedMessage
}
