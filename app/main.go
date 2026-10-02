package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/app/scanner"
	"github.com/codecrafters-io/interpreter-starter-go/app/token"
)

func main() {
	fmt.Fprintln(os.Stderr, "Logs from your program will appear here!")

	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" && command != "parse" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}
	// TODO refactor this later
	if command == "parse" {
		// needs refactoring
		filename := os.Args[2]
		fileContents, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
		s := scanner.New(string(fileContents))
		tokens, hadError := s.ScanTokens()
		// tester assuming that i just get true false nil
		for _, tok := range tokens {
			if tok.TokenType == token.TRUE || tok.TokenType == token.FALSE || tok.TokenType == token.NIL {
				fmt.Printf(tok.Lexeme)
			}
			if tok.TokenType == token.NUMBER || tok.TokenType == token.STRING {
				fmt.Printf("%v", tok.Literal)
			}
			if tok.TokenType == token.LEFTPAREN {
				fmt.Printf("%sgroup ", tok.Lexeme)
			}
			if tok.TokenType == token.IDENTIFIER {
				fmt.Printf("%s", tok.Literal)
			}
			if tok.TokenType == token.RIGHTPAREN {
				fmt.Printf("%s", tok.Lexeme)
			}
		}
		if hadError {
			os.Exit(65)
		}
		return
	}
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}
	s := scanner.New(string(fileContents))
	tokens, hadError := s.ScanTokens()
	for _, t := range tokens {
		fmt.Println(t.String())
	}

	if hadError {
		os.Exit(65)
	}
}
