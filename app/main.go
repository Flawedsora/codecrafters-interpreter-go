package main

import (
	"fmt"
	"os"
)

func lexer(textContent string) {
	// token_type lexeme literal
	// just handling parenthesis
	for _, r := range textContent {
		if r == '(' {
			fmt.Println("LEFT_PAREN ( null")
		} else if r == ')' {
			fmt.Println("RIGHT_PAREN ) null")
		}
	}
	// when converted to string EOF is removed
	fmt.Println("EOF  null")
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

	if len(fileContents) > 0 {
		textContent := string(fileContents)
		lexer(textContent)
	} else {
		fmt.Println("EOF  null") // Placeholder, replace this line when implementing the scanner
	}
}
