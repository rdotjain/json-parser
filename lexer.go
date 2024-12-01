package main

import (
	"unicode"
)

type Lexer struct {
	input  []rune
	pos    int
	length int
}

func NewLexer(input string) *Lexer {
	return &Lexer{input: []rune(input), length: len(input)}
}

func (l *Lexer) next() rune {
	if l.pos >= l.length {
		return -1
	}
	r := l.input[l.pos]
	l.pos++
	return r
}

func (l *Lexer) peek() rune {
	if l.pos >= l.length {
		return -1
	}
	return l.input[l.pos]
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.peek()) {
		l.next()
	}
}

func (l *Lexer) readString() string {
	var result []rune
	if l.next() != '"' {
		panic("Expected \" at the beginning of string")
	}
	for {
		r := l.next()
		if r == '"' {
			break
		}
		if r == -1 {
			panic("Unterminated string")
		}
		result = append(result, r)
	}
	return string(result)
}

func (l *Lexer) readNumber() string {
	var result []rune
	for unicode.IsDigit(l.peek()) || l.peek() == '.' { // 56 or 56.4
		result = append(result, l.next())
	}
	return string(result)
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()
	switch l.peek() {
	case '{':
		l.next()
		return Token{type_: TOKEN_LBRACE}
	case '}':
		l.next()
		return Token{type_: TOKEN_RBRACE}
	case '[':
		l.next()
		return Token{type_: TOKEN_LBRACKET}
	case ']':
		l.next()
		return Token{type_: TOKEN_RBRACKET}
	case ',':
		l.next()
		return Token{type_: TOKEN_COMMA}
	case ':':
		l.next()
		return Token{type_: TOKEN_COLON}
	case '"':
		return Token{type_: TOKEN_STRING, value: l.readString()}
	case 't':
		if l.matchKeyword("true") {
			return Token{type_: TOKEN_TRUE}
		}
	case 'f':
		if l.matchKeyword("false") {
			return Token{type_: TOKEN_FALSE}
		}
	case 'n':
		if l.matchKeyword("null") {
			return Token{type_: TOKEN_NULL}
		}
	default:
		if unicode.IsDigit(l.peek()) || l.peek() == '-' {
			return Token{type_: TOKEN_NUMBER, value: l.readNumber()}
		}
	}
	return Token{type_: TOKEN_EOF}
}

func (l *Lexer) matchKeyword(keyword string) bool {
	for _, r := range keyword {
		if l.next() != r {
			return false
		}
	}
	return true
}
