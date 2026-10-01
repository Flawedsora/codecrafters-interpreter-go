package token

import (
	"fmt"
)

type TokenType string

const (
	// Single-character tokens.
	LEFTPAREN  TokenType = "LEFT_PAREN"
	RIGHTPAREN TokenType = "RIGHT_PAREN"
	LEFTBRACE  TokenType = "LEFT_BRACE"
	RIGHTBRACE TokenType = "RIGHT_BRACE"
	COMMA      TokenType = "COMMA"
	DOT        TokenType = "DOT"
	MINUS      TokenType = "MINUS"
	PLUS       TokenType = "PLUS"
	SEMICOLON  TokenType = "SEMICOLON"
	SLASH      TokenType = "SLASH"
	STAR       TokenType = "STAR"
	// One or two character tokens.
	BANG         TokenType = "BANG"
	BANGEQUAL    TokenType = "BANG_EQUAL"
	EQUAL        TokenType = "EQUAL"
	EQUALEQUAL   TokenType = "EQUAL_EQUAL"
	GREATER      TokenType = "GREATER"
	GREATEREQUAL TokenType = "GREATER_EQUAL"
	LESS         TokenType = "LESS"
	LESSEQUAL    TokenType = "LESS_EQUAL"
	// Literals.
	IDENTIFIER TokenType = "IDENTIFIER"
	STRING     TokenType = "STRING"
	NUMBER     TokenType = "NUMBER"
	// Keywords
	AND    TokenType = "AND"
	CLASS  TokenType = "CLASS"
	ELSE   TokenType = "ELSE"
	FALSE  TokenType = "FALSE"
	FUN    TokenType = "FUN"
	FOR    TokenType = "FOR"
	IF     TokenType = "IF"
	NIL    TokenType = "NIL"
	OR     TokenType = "OR"
	PRINT  TokenType = "PRINT"
	RETURN TokenType = "RETURN"
	SUPER  TokenType = "SUPER"
	THIS   TokenType = "THIS"
	TRUE   TokenType = "TRUE"
	VAR    TokenType = "VAR"
	WHILE  TokenType = "WHILE"
	EOF    TokenType = "EOF"
)

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   any
	Line      int
}

func NewToken(tokentype TokenType, lexeme string, literal any, line int) Token {
	return Token{
		TokenType: tokentype,
		Lexeme:    lexeme,
		Literal:   literal,
		Line:      line,
	}
}

func (t Token) String() string {
	if t.Literal == nil {
		return fmt.Sprintf("%s %s null", t.TokenType, t.Lexeme)
	}
	return fmt.Sprintf("%s %s %v", t.TokenType, t.Lexeme, t.Literal)
}
