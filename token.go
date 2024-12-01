package main

type TokenType int

const (
	TOKEN_LBRACE TokenType = iota // '{'
	TOKEN_RBRACE                  // '}'
	TOKEN_LBRACKET                // '['
	TOKEN_RBRACKET                // ']'
	TOKEN_COMMA                   // ','
	TOKEN_COLON                   // ':'
	TOKEN_STRING                  // "string"
	TOKEN_NUMBER                  // number
	TOKEN_TRUE                    // true
	TOKEN_FALSE                   // false
	TOKEN_NULL                    // null
	TOKEN_EOF                     // end of input
)

type Token struct {
	type_ TokenType
	value string
}