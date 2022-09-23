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

func (p tokenParser_TestDouble) Minus() {
  p.Minus_CountOfCalls++
}

func (p tokenParser_TestDouble) Comma() {
  p.Comma_CountOfCalls++
}

func (p tokenParser_TestDouble) Plus() {
  p.Plus_CountOfCalls++
}

func (p tokenParser_TestDouble) Period() {
  p.Period_CountOfCalls++
}

func (p tokenParser_TestDouble) OpenSquareBracket() {
  p.OpenSquareBracket_CountOfCalls++
}

func (p tokenParser_TestDouble) CloseSquareBracket() {
  p.CloseSquareBracket_CountOfCalls++
}

func (p tokenParser_TestDouble) OpenAngleBracket() {
  p.OpenAngleBracket_CountOfCalls++
}

func (p tokenParser_TestDouble) CloseAngleBracket() {
  p.CloseAngleBracket_CountOfCalls++
}

func TestStuff(t *testing.T) {
  tokenParser := tokenParser_TestDouble{}
  lexer := parser.LexigraphicAnalyzer{ TokenParser: tokenParser }

  lexer.Analyze("")

  Equal(t, 1, tokenParser.CloseAngleBracket_CountOfCalls)
}
