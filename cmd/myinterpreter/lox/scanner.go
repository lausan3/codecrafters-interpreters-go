package lox

import (
	"fmt"
	"os"
)

type Scanner struct {
	FileContents string
}

func (s *Scanner) ScanTokens() {
	sawBadRune := false
	for _, char := range s.FileContents {
		switch char {
		case '(':
			fmt.Println("LEFT_PAREN ( null")
		case ')':
			fmt.Println("RIGHT_PAREN ) null")
		case '{':
			fmt.Println("LEFT_BRACE { null")
		case '}':
			fmt.Println("RIGHT_BRACE } null")
		case '.':
			fmt.Println("DOT . null")
		case ',':
			fmt.Println("COMMA , null")
		case '+':
			fmt.Println("PLUS + null")
		case '-':
			fmt.Println("MINUS - null")
		case ';':
			fmt.Println("SEMICOLON ; null")
		case '*':
			fmt.Println("STAR * null")
		default:
			fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %c\n", char)
			sawBadRune = true
		}
	}
	fmt.Println("EOF  null")

	if sawBadRune {
		os.Exit(65)
	}
}
