package main

// type Type string

// type Token struct {
// 	tokenType Type
// 	literal   string
// 	lineNo    int
// 	colNo     int
// }

var (
	EmptyLineByte          byte = byte('\n')
	WhiteSpaceByte         byte = byte(' ')
	ColonByte              byte = byte(':')
	CommaByte              byte = byte(',')
	LeftCurlyBraceByte     byte = byte('{')
	RightCurlyBraceByte    byte = byte('}')
	LeftSquareBracketByte  byte = byte('[')
	RightSquareBracketByte byte = byte(']')
	DoubleQuoteByte        byte = byte('"')
	BackSlashByte          byte = byte('\\')
	ForwardSlashByte       byte = byte('/')
)

var (
	LeftCurlyBrace   string = "{"
	RightCurlyBrace  string = "}"
	LeftSquareBrace  string = "["
	RightSquareBrace string = "]"
	Comma            string = ","
	Colon            string = ":"
	DoubleQuote      string = "\""
)

var BackSlashRune rune = rune('\\')

var (
	CheckSlice                  = []byte{LeftCurlyBraceByte, LeftSquareBracketByte, ColonByte, CommaByte, RightSquareBracketByte, RightCurlyBraceByte}
	AllowedCharsAfterEscapeChar = []rune{'"', '\\', '/', 'b', 'f', 'n', 'r', 't'}
)

// const (
// 	// Token/character we don't know about
// 	Illegal Type = "ILLEGAL"

// 	// End of file
// 	EOF Type = "EOF"

// 	// Literals
// 	String Type = "STRING"
// 	Number Type = "NUMBER"

// 	// The six structural tokens
// 	LeftBrace    Type = "{"
// 	RightBrace   Type = "}"
// 	LeftBracket  Type = "["
// 	RightBracket Type = "]"
// 	Comma        Type = ","
// 	Colon        Type = ":"

// 	// Values
// 	True  Type = "TRUE"
// 	False Type = "FALSE"
// 	Null  Type = "NULL"
// )
