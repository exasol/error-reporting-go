package exaerror_test

import (
	"github.com/exasol/error-reporting-go"
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
	renderedString := exaerror.New("E-TEST-1").Message("Something went wrong.").String()
	suite.Equal("E-TEST-1: Something went wrong.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithStringParameter() {
	renderedString := exaerror.New("E-TEST-2").Message("Unknown input {{input}}.").
		Parameter("input", "unknown").String()
	suite.Equal("E-TEST-2: Unknown input 'unknown'.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithParameterDescription() {
	renderedString := exaerror.New("E-TEST-2").Message("Unknown input {{input}}.").
		ParameterWithDescription("input", "unknown", "Input parameter.").String()
	suite.Equal("E-TEST-2: Unknown input 'unknown'.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithNilParameter() {
	renderedString := exaerror.New("E-TEST-2").Message("Unknown input {{input}}.").
		Parameter("input", nil).String()
	suite.Equal("E-TEST-2: Unknown input '<nil>'.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithIntParameter() {
	renderedString := exaerror.New("E-TEST-2").Message("Unknown input {{input}}.").
		Parameter("input", 10).String()
	suite.Equal("E-TEST-2: Unknown input '10'.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithMitigation() {
	renderedString := exaerror.New("E-TEST-2").Message("Too little disk space.").
		Mitigation("Delete something.").String()
	suite.Equal("E-TEST-2: Too little disk space. Delete something.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithUnquotedParameter() {
	renderedString := exaerror.New("E-TEST-2").Message("Unknown input {{input|uq}}.").
		Parameter("input", 2).String()
	suite.Equal("E-TEST-2: Unknown input 2.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithMissingParameter() {
	renderedString := exaerror.New("E-TEST-2").Message("Unknown input {{input}}.").String()
	suite.Equal("E-TEST-2: Unknown input {{input}}.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorCodeWithMissingParameterDefinition() {
	renderedString := exaerror.New("E-TEST-2").Message("Unknown input.").
		Parameter("input", 2).String()
	suite.Equal("E-TEST-2: Unknown input.", renderedString)
}

func (suite *ExaErrorTestSuite) TestShouldImplementErrorInterface() {
	err := exaerror.New("E-TEST-2").Message("Unknown input.").
		Parameter("input", 2)

	suite.Error(err, "Error should not be nil")
	suite.EqualError(err, "E-TEST-2: Unknown input.")

}

func (suite *ExaErrorTestSuite) TestErrorMessagefCodeWithStringParameter() {
	renderedString := exaerror.New("E-TEST-2").Messagef("Unknown input {{input}}.", "Value").String()
	suite.Equal("E-TEST-2: Unknown input 'Value'.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorMessagefCodeWithMissingParameter() {
	renderedString := exaerror.New("E-TEST-2").Messagef("Unknown input {{input}} {{input-missing}}.", "Value").String()
	suite.Equal("E-TEST-2: Unknown input 'Value' {{input-missing}}.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorMessagefCodeWithTooManyParameter() {
	renderedString := exaerror.New("E-TEST-2").Messagef("Unknown input {{input}}.", "Value", "Value2").String()
	suite.Equal("E-TEST-2: Unknown input 'Value'.", renderedString)
}

func (suite *ExaErrorTestSuite) TestErrorMessagefCodeWithMultipleParameter() {
	renderedString := exaerror.New("E-TEST-2").Messagef("Unknown input {{input}} and {{code}}.", "Value", 42).String()
	suite.Equal("E-TEST-2: Unknown input 'Value' and '42'.", renderedString)
}



