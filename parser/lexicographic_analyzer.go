package parser

type tokenParser interface {
  CloseAngleBracket()
  OpenAngleBracket()
  Plus()
  Minus()
  Period()
  Comma()
  OpenSquareBracket()
  CloseSquareBracket()
}

type LexigraphicAnalyzer struct {
  TokenParser tokenParser
}

func (l LexigraphicAnalyzer) Analyze(input string) {
  l.TokenParser.CloseAngleBracket()
}
