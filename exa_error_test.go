package error_reporting_go

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type ExaErrorTestSuite struct {
	suite.Suite
}

func TestExaErrorSuite(t *testing.T) {
	suite.Run(t, new(ExaErrorTestSuite))
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithMessage() {
	renderedString := ExaError("E-TEST-1").Message("Something went wrong.").String()
	suite.Equal("E-TEST-1: Something went wrong.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithStringParameter() {
	renderedString := ExaError("E-TEST-2").Message("Unknown input {{input}}.").
		Parameter("input", "unknown").String()
	suite.Equal("E-TEST-2: Unknown input 'unknown'.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithParameterDescription() {
	renderedString := ExaError("E-TEST-2").Message("Unknown input {{input}}.").
		ParameterWithDescription("input", "unknown", "Input parameter.").String()
	suite.Equal("E-TEST-2: Unknown input 'unknown'.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithNilParameter() {
	renderedString := ExaError("E-TEST-2").Message("Unknown input {{input}}.").
		Parameter("input", nil).String()
	suite.Equal("E-TEST-2: Unknown input '<nil>'.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithIntParameter() {
	renderedString := ExaError("E-TEST-2").Message("Unknown input {{input}}.").
		Parameter("input", 10).String()
	suite.Equal("E-TEST-2: Unknown input '10'.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithMitigation() {
	renderedString := ExaError("E-TEST-2").Message("Too little disk space.").
		Mitigation("Delete something.").String()
	suite.Equal("E-TEST-2: Too little disk space. Delete something.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithUnquotedParameter() {
	renderedString := ExaError("E-TEST-2").Message("Unknown input {{input|uq}}.").
		Parameter("input", 2).String()
	suite.Equal("E-TEST-2: Unknown input 2.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithMissingParameter() {
	renderedString := ExaError("E-TEST-2").Message("Unknown input {{input}}.").String()
	suite.Equal("E-TEST-2: Unknown input {{input}}.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithMissingParameterDefinition() {
	renderedString := ExaError("E-TEST-2").Message("Unknown input.").
		Parameter("input", 2).String()
	suite.Equal("E-TEST-2: Unknown input.", renderedString)
}
