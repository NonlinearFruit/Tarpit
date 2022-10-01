package parser_test

import (
  "testing"
  . "github.com/stretchr/testify/assert"
  "github.com/nonlinearfruit/tarpit/parser"
)

type tokenParser_TestDouble struct {
  CloseAngleBracket_CountOfCalls int
  OpenAngleBracket_CountOfCalls int
  Plus_CountOfCalls int
  Minus_CountOfCalls int
  Period_CountOfCalls int
  Comma_CountOfCalls int
  OpenSquareBracket_CountOfCalls int
  CloseSquareBracket_CountOfCalls int
}

func (p *tokenParser_TestDouble) Minus() {
  p.Minus_CountOfCalls++
}

func (p *tokenParser_TestDouble) Comma() {
  p.Comma_CountOfCalls++
}

func (p *tokenParser_TestDouble) Plus() {
  p.Plus_CountOfCalls++
}

func (p *tokenParser_TestDouble) Period() {
  p.Period_CountOfCalls++
}

func (p *tokenParser_TestDouble) OpenSquareBracket() {
  p.OpenSquareBracket_CountOfCalls++
}

func (p *tokenParser_TestDouble) CloseSquareBracket() {
  p.CloseSquareBracket_CountOfCalls++
}

func (p *tokenParser_TestDouble) OpenAngleBracket() {
  p.OpenAngleBracket_CountOfCalls++
}

func (p *tokenParser_TestDouble) CloseAngleBracket() {
  p.CloseAngleBracket_CountOfCalls = p.CloseAngleBracket_CountOfCalls + 1
}

func TestCloseAngleBracket(t *testing.T) {
  tokenParser := tokenParser_TestDouble{}
  lexer := parser.LexigraphicAnalyzer{ TokenParser: &tokenParser }

  lexer.Analyze(">")

  Equal(t, 1, tokenParser.CloseAngleBracket_CountOfCalls)
}

func TestOpenAngleBracket(t *testing.T) {
  tokenParser := tokenParser_TestDouble{}
  lexer := parser.LexigraphicAnalyzer{ TokenParser: &tokenParser }

  lexer.Analyze("<")

  Equal(t, 1, tokenParser.OpenAngleBracket_CountOfCalls)
}

func TestOpenSquareBracket(t *testing.T) {
  tokenParser := tokenParser_TestDouble{}
  lexer := parser.LexigraphicAnalyzer{ TokenParser: &tokenParser }

  lexer.Analyze("[")

  Equal(t, 1, tokenParser.OpenSquareBracket_CountOfCalls)
}

func TestCloseSquareBracket(t *testing.T) {
  tokenParser := tokenParser_TestDouble{}
  lexer := parser.LexigraphicAnalyzer{ TokenParser: &tokenParser }

  lexer.Analyze("]")

  Equal(t, 1, tokenParser.CloseSquareBracket_CountOfCalls)
}

func TestComma(t *testing.T) {
  tokenParser := tokenParser_TestDouble{}
  lexer := parser.LexigraphicAnalyzer{ TokenParser: &tokenParser }

  lexer.Analyze(",")

  Equal(t, 1, tokenParser.Comma_CountOfCalls)
}

func TestPeriod(t *testing.T) {
  tokenParser := tokenParser_TestDouble{}
  lexer := parser.LexigraphicAnalyzer{ TokenParser: &tokenParser }

  lexer.Analyze(".")

  Equal(t, 1, tokenParser.Period_CountOfCalls)
}

func TestPlus(t *testing.T) {
  tokenParser := tokenParser_TestDouble{}
  lexer := parser.LexigraphicAnalyzer{ TokenParser: &tokenParser }

  lexer.Analyze("+")

  Equal(t, 1, tokenParser.Plus_CountOfCalls)
}

func TestMinus(t *testing.T) {
  tokenParser := tokenParser_TestDouble{}
  lexer := parser.LexigraphicAnalyzer{ TokenParser: &tokenParser }

  lexer.Analyze("-")

  Equal(t, 1, tokenParser.Minus_CountOfCalls)
}

func TestParsersMultipleCharacters(t *testing.T) {
  tokenParser := tokenParser_TestDouble{}
  lexer := parser.LexigraphicAnalyzer{ TokenParser: &tokenParser }

  lexer.Analyze("---")

  Equal(t, 3, tokenParser.Minus_CountOfCalls)
}
