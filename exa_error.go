package error_reporting_go

import (
	"fmt"
	"strings"
)

type ErrorMessageBuilder struct {
	errorCode  string
	message    string
	parameters []parameter
	mitigation string
}

type parameter struct {
	name  string
	value string
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

func (builder *ErrorMessageBuilder) Parameter(name string, value interface{}) *ErrorMessageBuilder {
	builder.parameters = append(builder.parameters, parameter{name, fmt.Sprintf("%v", value)})
	return builder
}

func (builder *ErrorMessageBuilder) ParameterWithDescription(name string, value interface{}, description string) *ErrorMessageBuilder {
	builder.Parameter(name, value)
	return builder
}

func (builder *ErrorMessageBuilder) Mitigation(mitigation string) *ErrorMessageBuilder {
	builder.mitigation = mitigation
	return builder
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

func (builder *ErrorMessageBuilder) Error() string  {
    return builder.String()
}

func formatMessage(builder *ErrorMessageBuilder) string {
	var formattedMessage = builder.message
	for _, parameter := range builder.parameters {
		formattedMessage = strings.Replace(formattedMessage, "{{"+parameter.name+"}}", "'"+parameter.value+"'", -1)
		formattedMessage = strings.Replace(formattedMessage, "{{"+parameter.name+"|uq}}", parameter.value, -1)
	}
	return formattedMessage
}
