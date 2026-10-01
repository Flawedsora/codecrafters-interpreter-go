package scanner

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/app/token"
)

type Scanner struct {
	source            string
	tokens            []token.Token
	start, curr, line int
	hadError          bool
}

func New(source string) *Scanner {
	return &Scanner{
		source: source,
		line:   1,
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.curr >= len(s.source)
}

func (s *Scanner) advance() byte {
	// move forward
	c := s.source[s.curr]
	s.curr++
	return c
}

func (s *Scanner) peek() byte {
	// give current token
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.curr]
}

func (s *Scanner) peekNext() byte {
	// useful later
	if s.curr+1 >= len(s.source) {
		return 0
	}
	return s.source[s.curr+1]
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.source[s.curr] != expected {
		return false
	}
	s.curr++
	return true
}

func (s *Scanner) addToken(t token.TokenType) {
	text := s.source[s.start:s.curr]
	s.tokens = append(s.tokens, token.NewToken(t, text, nil, s.line))
}

func (s *Scanner) addTokenLiteral(t token.TokenType, literal any) {
	text := s.source[s.start:s.curr]
	s.tokens = append(s.tokens, token.NewToken(t, text, literal, s.line))
}

func (s *Scanner) ScanTokens() ([]token.Token, bool) {
	for !s.isAtEnd() {
		s.start = s.curr
		s.scanToken()
	}
	s.tokens = append(s.tokens, token.NewToken(token.EOF, "", nil, s.line))
	return s.tokens, s.hadError
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		s.advance()
	}
	if s.isAtEnd() {
		s.hadError = true
		fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", s.line)
		return
	}
	s.advance()
	val := s.source[s.start+1 : s.curr-1]
	s.addTokenLiteral(token.STRING, val)
}

func (s *Scanner) scanToken() {
	c := s.advance()
	if isAlpha(c) {
		s.identifier()
		return
	}
	if isDigit(c) {
		s.number()
		return
	}
	if c == '"' {
		s.string()
		return
	}
	// get current value and store it move idx forward also
	switch c {
	case '(':
		s.addToken(token.LEFTPAREN)
	case ')':
		s.addToken(token.RIGHTPAREN)
	case '{':
		s.addToken(token.LEFTBRACE)
	case '}':
		s.addToken(token.RIGHTBRACE)
	case ',':
		s.addToken(token.COMMA)
	case '.':
		s.addToken(token.DOT)
	case '-':
		s.addToken(token.MINUS)
	case '+':
		s.addToken(token.PLUS)
	case ';':
		s.addToken(token.SEMICOLON)
	case '*':
		s.addToken(token.STAR)
	case '!':
		if s.match('=') {
			// if match found then in s.match() increment happens
			s.addToken(token.BANGEQUAL)
		} else {
			s.addToken(token.BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(token.EQUALEQUAL)
		} else {
			s.addToken(token.EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(token.LESSEQUAL)
		} else {
			s.addToken(token.LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(token.GREATEREQUAL)
		} else {
			s.addToken(token.GREATER)
		}
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(token.SLASH)
		}
	case ' ', '\r', '\t':
		// ignore
	case '\n':
		s.line++
	default:
		s.hadError = true
		fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", s.line, c)
	}
}
