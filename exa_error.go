package exaerror

import (
	"fmt"
	"regexp"
	"strings"
)

type ExaError struct {
	errorCode       string
	message         string
	parameters      []parameter
	mitigations     []string
	mitigationCount int
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
			paramName := strings.TrimSuffix(paramNames[i][1], "|uq")
			_ = builder.Parameter(paramName, param)
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
	builder.mitigationCount += 1
	builder.mitigations = append(builder.mitigations, mitigation)
	return builder
}

func (builder *ExaError) String() string {
	var stringBuilder strings.Builder
	formattedMessage := formatMessage(builder)
	stringBuilder.WriteString(fmt.Sprintf("%s: %s", builder.errorCode, formattedMessage))

	if len(builder.mitigations) > 0 {
		var mitigationString string
		if len(builder.mitigations) == 1 {
			mitigationString = ""
			for _, mitigation := range builder.mitigations {
				mitigationString += mitigation
			}
		} else {
			mitigationString = "Known mitigations:"
			for _, mitigation := range builder.mitigations {
				mitigationString += "\n" + "* " + mitigation
			}
		}
		formattedMitigation := formatMitigation(builder, mitigationString)
		stringBuilder.WriteString(fmt.Sprintf(" %s", formattedMitigation))
	}

	return stringBuilder.String()
}

func (builder *ExaError) Error() string {
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
func formatMitigation(builder *ExaError, mitigationString string) string {
	var formattedMitigation = mitigationString
	for _, parameter := range builder.parameters {
		formattedMitigation = strings.Replace(formattedMitigation, "{{"+parameter.name+"}}", "'"+parameter.value+"'", -1)
		formattedMitigation = strings.Replace(formattedMitigation, "{{"+parameter.name+"|uq}}", parameter.value, -1)
	}
	return formattedMitigation
}
