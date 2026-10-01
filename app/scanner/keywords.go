package scanner

import "github.com/codecrafters-io/interpreter-starter-go/app/token"

var Keywords = map[string]token.TokenType{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"false":  token.FALSE,
	"for":    token.FOR,
	"fun":    token.FUN,
	"if":     token.IF,
	"nil":    token.NIL,
	"or":     token.OR,
	"print":  token.PRINT,
	"return": token.RETURN,
	"super":  token.SUPER,
	"this":   token.THIS,
	"true":   token.TRUE,
	"var":    token.VAR,
	"while":  token.WHILE,
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.curr]
	if typ, ok := Keywords[text]; ok {
		s.addToken(typ)
	} else {
		s.addToken(token.IDENTIFIER)
	}
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}
	// only 21.23 allowed 23.a not
	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance() // eat '.'
		for isDigit(s.peek()) {
			s.advance()
		}
	}
	text := s.source[s.start:s.curr]
	s.addTokenLiteral(token.NUMBER, numberLiteral(text))
}

func numberLiteral(text string) string {
	// copy of old numTransform in main.go:33-43
	// "123.4500" -> "123.45", "123.00" -> "123.0", "123" -> "123.0"
	if len(text) == 0 {
		return "0.0"
	}
	hasDot := false
	for i := 0; i < len(text); i++ {
		if text[i] == '.' {
			hasDot = true
			break
		}
	}
	if !hasDot {
		return text + ".0"
	}
	idx := len(text) - 1
	for idx >= 0 && text[idx] == '0' {
		idx--
	}
	if text[idx] == '.' {
		return text[0:idx+1] + "0"
	}
	return text[0 : idx+1]
}
