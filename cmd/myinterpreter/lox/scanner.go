package lox

import (
	"fmt"
	"os"
)

type scanner struct {
	Source  string
	Current int
	Start   int
	Line    int
}

func NewScanner(source string) *scanner {
	return &scanner{Source: source}
}

// Scan the source string for its tokens.
func (s *scanner) ScanTokens() {
	sawBadRune := false
	s.Current = 0
	s.Line = 1

	for !s.isAtEnd() {
		s.Start = s.Current
		if !s.scanToken() {
			sawBadRune = true
		}
	}

	fmt.Println("EOF  null")

	if sawBadRune {
		os.Exit(65)
	}
}

// Scan the current token in the source string and return whether we recognized it.
func (s *scanner) scanToken() (ok bool) {
	char := s.advance()

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
	case '=':
		if s.match('=') {
			fmt.Println("EQUAL_EQUAL == null")
		} else {
			fmt.Println("EQUAL = null")
		}
	case '!':
		if s.match('=') {
			fmt.Println("BANG_EQUAL != null")
		} else {
			fmt.Println("BANG ! null")
		}
	case '<':
		if s.match('=') {
			fmt.Println("LESS_EQUAL <= null")
		} else {
			fmt.Println("LESS < null")
		}
	case '>':
		if s.match('=') {
			fmt.Println("GREATER_EQUAL >= null")
		} else {
			fmt.Println("GREATER > null")
		}
	case '/':
		// Ignore comments
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
				// fmt.Printf("skipping comment %c\n", chr)
			}
		} else {
			fmt.Println("SLASH / null")
		}
	case '\t', ' ':
	case '\n':
		s.Line++
	default:
		fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", s.Line, char)
		return false
	}

	return true
}

// Return the rune at the current pointer and advance the pointer to the next index.
func (s *scanner) advance() rune {
	currentRune := rune(s.Source[s.Current])
	s.Current += 1
	return currentRune
}

// Match the next rune with an expected one.
func (s *scanner) match(expected rune) bool {
	if s.isAtEnd() || s.Source[s.Current] != byte(expected) {
		return false
	}

	s.Current += 1
	return true
}

// Checks if the current pointer is not at the end of source string.
func (s *scanner) isAtEnd() bool {
	return s.Current >= len(s.Source)
}

// Return the next rune in the source string, else return a space.
func (s *scanner) peek() rune {
	if s.isAtEnd() {
		return ' '
	}

	return rune(s.Source[s.Current])
}
