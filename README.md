# json-parser

### Example Usage
``` bash
go run main.go myfile.json
```


### Lexer:

`NewLexer(input string) *Lexer`
Creates a new Lexer for the provided input string.

`(*Lexer) next()` rune
Advances to the next character in the input and returns it. Returns -1 if the end is reached.

`(*Lexer) peek()` rune
Returns the current character without advancing the position. Returns -1 at the end of input.

`(*Lexer) skipWhitespace()`
Skips over any whitespace characters in the input.

`(*Lexer) readString()` string
Reads a string token enclosed in double quotes. Panics if the string is unterminated.

`(*Lexer) readNumber()` string
Reads a number (integer or float) from the input.

`(*Lexer) NextToken()` Token
Returns the next token based on the current input. Skips whitespace and handles different token types (e.g., braces, strings, numbers).

### Parser:

`NewParser(input string) *Parser`
Creates and initializes a new Parser.

`(*Parser) advance()`
Advances to the next token.

`(*Parser) parseValue()` bool
Parses a single value (string, number, boolean, null, object, or array). Returns true if successful, otherwise false.

`(*Parser) parseObject()` bool
Parses an object ({key: value}). Returns true if successful, otherwise false.

`(*Parser) parseArray()` bool
Parses an array ([value1, value2, ...]). Returns true if successful, otherwise false