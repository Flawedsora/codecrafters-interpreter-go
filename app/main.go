package main

import (
	"fmt"
	"os"
)

type Token struct {
	Type, Lexeme, Literal string
}

func lexer(textContent string) []Token {
	// token_type lexeme literal
	tokens := []Token{}
	for i := 0; i < len(textContent); i++ {
		switch c := textContent[i]; c {
		case '(':
			tokens = append(tokens, Token{"LEFT_PAREN", "(", "null"})
		case ')':
			tokens = append(tokens, Token{"RIGHT_PAREN", ")", "null"})
		case '{':
			tokens = append(tokens, Token{"LEFT_BRACE", "{", "null"})
		case '}':
			tokens = append(tokens, Token{"RIGHT_BRACE", "}", "null"})
		case ',':
			tokens = append(tokens, Token{"COMMA", ",", "null"})
		case '.':
			tokens = append(tokens, Token{"DOT", ".", "null"})
		case '-':
			tokens = append(tokens, Token{"MINUS", "-", "null"})
		case '+':
			tokens = append(tokens, Token{"PLUS", "+", "null"})
		case ';':
			tokens = append(tokens, Token{"SEMICOLON", ";", "null"})
		case '/':
			tokens = append(tokens, Token{"SLASH", "/", "null"})
		case '*':
			tokens = append(tokens, Token{"STAR", "*", "null"})
		}
	}
	tokens = append(tokens, Token{"EOF", "", "null"})
	return tokens
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
	tokens := lexer(string(fileContents))
	for _, t := range tokens {
		fmt.Printf("%s %s %s\n", t.Type, t.Lexeme, t.Literal)
	}
}
