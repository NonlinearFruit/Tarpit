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

func (l *LexigraphicAnalyzer) Analyze(input string) {
  for _, c := range input {
    if c == '>' {
      l.TokenParser.CloseAngleBracket()
    } else if c == '<' {
      l.TokenParser.OpenAngleBracket()
    } else if c == '[' {
      l.TokenParser.OpenSquareBracket()
    } else if c == ']' {
      l.TokenParser.CloseSquareBracket()
    } else if c == ',' {
      l.TokenParser.Comma()
    } else if c == '.' {
      l.TokenParser.Period()
    } else if c == '+' {
      l.TokenParser.Plus()
    } else if c == '-' {
      l.TokenParser.Minus()
    } else {
      // Comment character
    }
  }
}
