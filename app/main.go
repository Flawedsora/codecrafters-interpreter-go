package main

import (
	"fmt"
	"os"
)

type Token struct {
	Type, Lexeme, Literal string
}

func peek(textContent string, i int) byte {
	if i+1 < len(textContent) {
		return textContent[i+1]
	}
	return 0
}

func lexer(textContent string) ([]Token, bool) {
	// token_type lexeme literal
	tokens := []Token{}
	hadError := false
	tlen := len(textContent)
	for i := 0; i < tlen; {
		c := textContent[i]
		if c == ' ' || c == '\t' || c == '\r' || c == '\n' {
			i += 1
			continue
		}
		if i+1 < tlen && c == '/' && textContent[i+1] == '/' {
			for i < tlen && textContent[i] != '\n' {
				i += 1
			}
			if i < tlen && textContent[i] == '\n' {
				i += 1
			}
			continue
		}
		if i >= tlen {
			break
		}
		switch c {
		case '(':
			tokens = append(tokens, Token{"LEFT_PAREN", "(", "null"})
			i += 1
		case ')':
			tokens = append(tokens, Token{"RIGHT_PAREN", ")", "null"})
			i += 1
		case '{':
			tokens = append(tokens, Token{"LEFT_BRACE", "{", "null"})
			i += 1
		case '}':
			tokens = append(tokens, Token{"RIGHT_BRACE", "}", "null"})
			i += 1
		case ',':
			tokens = append(tokens, Token{"COMMA", ",", "null"})
			i += 1
		case '.':
			tokens = append(tokens, Token{"DOT", ".", "null"})
			i += 1
		case '-':
			tokens = append(tokens, Token{"MINUS", "-", "null"})
			i += 1
		case '+':
			tokens = append(tokens, Token{"PLUS", "+", "null"})
			i += 1
		case ';':
			tokens = append(tokens, Token{"SEMICOLON", ";", "null"})
			i += 1
		case '/':
			if peek(textContent, i) == '/' {
				// TODO : refactor
				for i < len(textContent) && textContent[i] != '\n' {
					i += 1
				}
			} else {
				tokens = append(tokens, Token{"SLASH", "/", "null"})
				i += 1
			}
		case '*':
			tokens = append(tokens, Token{"STAR", "*", "null"})
			i += 1

		case '=':
			next := peek(textContent, i)
			if next == '=' {
				tokens = append(tokens, Token{"EQUAL_EQUAL", "==", "null"})
				i += 2
			} else {
				tokens = append(tokens, Token{"EQUAL", "=", "null"})
				i += 1
			}
		case '!':
			next := peek(textContent, i)
			if next == '=' {
				tokens = append(tokens, Token{"BANG_EQUAL", "!=", "null"})
				i += 2
			} else {
				tokens = append(tokens, Token{"BANG", "!", "null"})
				i += 1
			}
		case '<':
			next := peek(textContent, i)
			if next == '=' {
				tokens = append(tokens, Token{"LESS_EQUAL", "<=", "null"})
				i += 2
			} else {
				tokens = append(tokens, Token{"LESS", "<", "null"})
				i += 1
			}
		case '>':
			next := peek(textContent, i)
			if next == '=' {
				tokens = append(tokens, Token{"GREATER_EQUAL", ">=", "null"})
				i += 2
			} else {
				tokens = append(tokens, Token{"GREATER", ">", "null"})
				i += 1
			}
		default:
			hadError = true
			fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %c\n", c) // currently just one line
			i += 1
		}
	}
	tokens = append(tokens, Token{"EOF", "", "null"})
	return tokens, hadError
}

func main() {
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	tokens, hadError := lexer(string(fileContents))
	for _, t := range tokens {
		fmt.Printf("%s %s %s\n", t.Type, t.Lexeme, t.Literal)
	}
	if hadError {
		os.Exit(65)
	}
}
