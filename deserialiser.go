package main

func Deserializer(content []byte) bool {
	tokens := Lexer(content)
	_, result := Parser(tokens)
	return result
}