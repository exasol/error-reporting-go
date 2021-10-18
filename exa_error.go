package exaerror

import (
	"fmt"
	"regexp"
	"strings"
)



type ExaError struct {
	errorCode  string
	message    string
	parameters []parameter
	mitigation string
}

type parameter struct {
	name  string
	value string
}

func New(errorCode string) *ExaError {
	return &ExaError{
		errorCode: errorCode,
	}
}

func (builder *ExaError) Message(message string) *ExaError {
	builder.message = message
	return builder
}

func (builder *ExaError) Messagef(format string, a ...interface{}) *ExaError {
	builder.message = format
	rex := regexp.MustCompile(`{{(.*?)}}`)
	paramNames := rex.FindAllStringSubmatch(format, -1)

	for i, param := range a {
		if i < len(paramNames) {
			_ = builder.Parameter(paramNames[i][1], param)
		}
	}
	return builder
}

func (builder *ExaError) Parameter(name string, value interface{}) *ExaError {
	builder.parameters = append(builder.parameters, parameter{name, fmt.Sprintf("%v", value)})
	return builder
}

func (builder *ExaError) ParameterWithDescription(name string, value interface{}, description string) *ExaError {
	return builder.Parameter(name, value)
}

func (builder *ExaError) Mitigation(mitigation string) *ExaError {
	builder.mitigation = mitigation
	return builder
}

func (builder *ExaError) String() string {
	var stringBuilder strings.Builder
	formattedMessage := formatMessage(builder)
	stringBuilder.WriteString(fmt.Sprintf("%s: %s", builder.errorCode, formattedMessage))
	if len(builder.mitigation) > 0 {
		stringBuilder.WriteString(fmt.Sprintf(" %s", builder.mitigation))
	}
	return stringBuilder.String()
}

func (builder *ExaError) Error() string  {
    return builder.String()
}

func formatMessage(builder *ExaError) string {
	var formattedMessage = builder.message
	for _, parameter := range builder.parameters {
		formattedMessage = strings.Replace(formattedMessage, "{{"+parameter.name+"}}", "'"+parameter.value+"'", -1)
		formattedMessage = strings.Replace(formattedMessage, "{{"+parameter.name+"|uq}}", parameter.value, -1)
	}
	return formattedMessage
}
