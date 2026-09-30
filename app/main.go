package main

import (
	"fmt"
	"os"
	"strconv"
	"unicode"
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

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c rune) bool {
	return unicode.IsLetter(c) || c == '_'
}

func isAlphaNumeric(c rune) bool {
	return isAlpha(c) || unicode.IsDigit(c)
}

func numTransform(numberpart string) string {
	// trimming logic just for decimals
	idx := len(numberpart) - 1
	for idx >= 0 && numberpart[idx] == '0' {
		idx--
	}
	if numberpart[idx] == '.' {
		return numberpart[0:idx+1] + "0"
	}
	return numberpart[0 : idx+1]
}

func reservedWords(textToken string, i int) (Token, int) {
	// taking input string and location in string
	// Identifier part + next char which can't be part of identifier
	if textToken[i] == 'a' {
		if peek(textToken, i) == 'n' && peek(textToken, i+1) == 'd' && !isAlphaNumeric(rune(peek(textToken, i+2))) {
			return Token{"AND", "and", "null"}, i + 3
		}
	} else if textToken[i] == 'c' {
		if peek(textToken, i) == 'l' && peek(textToken, i+1) == 'a' && peek(textToken, i+2) == 's' && peek(textToken, i+3) == 's' && !isAlphaNumeric(rune(peek(textToken, i+4))) {
			return Token{"CLASS", "class", "null"}, i + 5
		}
	} else if textToken[i] == 'e' {
		if peek(textToken, i) == 'l' && peek(textToken, i+1) == 's' && peek(textToken, i+2) == 'e' && !isAlphaNumeric(rune(peek(textToken, i+3))) {
			return Token{"ELSE", "else", "null"}, i + 4
		}
	} else if textToken[i] == 'f' {
		if peek(textToken, i) == 'a' {
			if peek(textToken, i+1) == 'l' && peek(textToken, i+2) == 's' && peek(textToken, i+3) == 'e' && !isAlphaNumeric(rune(peek(textToken, i+4))) {
				return Token{"FALSE", "false", "null"}, i + 5
			}
		} else if peek(textToken, i) == 'o' {
			if peek(textToken, i+1) == 'r' && !isAlphaNumeric(rune(peek(textToken, i+2))) {
				return Token{"FOR", "for", "null"}, i + 3
			}
		} else if peek(textToken, i) == 'u' {
			if peek(textToken, i+1) == 'n' && !isAlphaNumeric(rune(peek(textToken, i+2))) {
				return Token{"FUN", "fun", "null"}, i + 4
			}
		}
	} else if textToken[i] == 'i' {
		if peek(textToken, i) == 'f' && !isAlphaNumeric(rune(peek(textToken, i+1))) {
			return Token{"IF", "if", "null"}, i + 2
		}
	} else if textToken[i] == 'n' {
		if peek(textToken, i) == 'i' && peek(textToken, i+1) == 'l' && !isAlphaNumeric(rune(peek(textToken, i+2))) {
			return Token{"NIL", "nil", "null"}, i + 3
		}
	} else if textToken[i] == 'o' {
		if peek(textToken, i) == 'r' && !isAlphaNumeric(rune(peek(textToken, i+1))) {
			return Token{"OR", "or", "null"}, i + 2
		}
	} else if textToken[i] == 'p' {
		if peek(textToken, i) == 'r' && peek(textToken, i+1) == 'i' && peek(textToken, i+2) == 'n' && peek(textToken, i+3) == 't' && !isAlphaNumeric(rune(peek(textToken, i+4))) {
			return Token{"PRINT", "print", "null"}, i + 5
		}
	} else if textToken[i] == 'r' {
		if peek(textToken, i) == 'e' && peek(textToken, i+1) == 't' && peek(textToken, i+2) == 'u' && peek(textToken, i+3) == 'r' && peek(textToken, i+4) == 'n' && !isAlphaNumeric(rune(peek(textToken, i+5))) {
			return Token{"RETURN", "return", "null"}, i + 6
		}
	} else if textToken[i] == 's' {
		if peek(textToken, i) == 'u' && peek(textToken, i+1) == 'p' && peek(textToken, i+2) == 'e' && peek(textToken, i+3) == 'r' && !isAlphaNumeric(rune(peek(textToken, i+4))) {
			return Token{"SUPER", "super", "null"}, i + 5
		}
	} else if textToken[i] == 't' {
		if peek(textToken, i) == 'h' {
			if peek(textToken, i+1) == 'i' && peek(textToken, i+2) == 's' && !isAlphaNumeric(rune(peek(textToken, i+3))) {
				return Token{"THIS", "this", "null"}, i + 4
			}
		} else if peek(textToken, i) == 'r' {
			if peek(textToken, i+1) == 'u' && peek(textToken, i+2) == 'e' && !isAlphaNumeric(rune(peek(textToken, i+3))) {
				return Token{"TRUE", "true", "null"}, i + 4
			}
		}
	} else if textToken[i] == 'v' {
		if peek(textToken, i) == 'a' && peek(textToken, i+1) == 'r' && !isAlphaNumeric(rune(peek(textToken, i+2))) {
			return Token{"VAR", "var", "null"}, i + 3
		}
	} else if textToken[i] == 'w' {
		if peek(textToken, i) == 'h' && peek(textToken, i+1) == 'i' && peek(textToken, i+2) == 'l' && peek(textToken, i+3) == 'e' && !isAlphaNumeric(rune(peek(textToken, i+4))) {
			return Token{"WHILE", "while", "null"}, i + 5
		}
	}
	return Token{"", "", ""}, i
}

func lexer(textContent string) ([]Token, bool) {
	// token_type lexeme literal
	tokens := []Token{}
	hadError := false
	tlen := len(textContent)
	lcnt := 1
	for i := 0; i < tlen; {
		c := textContent[i]
		tk, nval := reservedWords(textContent, i)
		if i != nval {
			tokens = append(tokens, tk)
			i = nval
			continue
		}
		if isAlpha(rune(c)) {
			start := i
			for i < tlen && isAlphaNumeric(rune(textContent[i])) {
				i += 1
			}
			end := i
			variablePart := textContent[start:end]
			tokens = append(tokens, Token{"IDENTIFIER", variablePart, "null"})
			continue
		}
		if isDigit(c) {
			start := i
			cntDecimal := 0
			for i < tlen && (isDigit(textContent[i]) || textContent[i] == '.') {
				if cntDecimal > 1 {
					break
				}
				if textContent[i] == '.' {
					cntDecimal += 1
				}
				i++
			}
			fmt.Fprintf(os.Stderr, "Value of i is %d", i)
			numberPart := textContent[start:i]
			var literalPart string
			if cntDecimal == 0 {
				literalPart = numberPart + ".0"
			} else {
				literalPart = numTransform(numberPart)
			}
			// start to end inclusive is needed
			tokens = append(tokens, Token{"NUMBER", numberPart, literalPart})
			continue
		}
		if c == ' ' || c == '\t' || c == '\r' || c == '\n' {
			if c == '\n' {
				lcnt += 1
			}
			i += 1
			continue
		}
		if i+1 < tlen && c == '/' && textContent[i+1] == '/' {
			for i < tlen && textContent[i] != '\n' {
				i += 1
			}
			if i < tlen && textContent[i] == '\n' {
				lcnt += 1
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
		case '*':
			tokens = append(tokens, Token{"STAR", "*", "null"})
			i += 1
		case '/':
			tokens = append(tokens, Token{"SLASH", "/", "null"})
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
		case '"':
			start := i
			i++
			for i < tlen && textContent[i] != '"' {
				i++
			}
			if i >= tlen || textContent[i] != '"' {
				hadError = true
				fmt.Fprintf(os.Stderr, "[line %s] Error: Unterminated string.", strconv.Itoa(lcnt))
				i++
				continue
			}
			stringPart := textContent[start : i+1]
			literalPart := textContent[start+1 : i]
			tokens = append(tokens, Token{"STRING", stringPart, literalPart})
			i++
		default:
			hadError = true
			fmt.Fprintf(os.Stderr, "[line %s] Error: Unexpected character: %c\n", strconv.Itoa(lcnt), c) // currently just one line
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
