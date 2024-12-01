package main

type Parser struct {
	lexer *Lexer
	token Token
}

func NewParser(input string) *Parser {
	lexer := NewLexer(input)
	return &Parser{lexer: lexer, token: lexer.NextToken()}
}

func (p *Parser) advance() {
	p.token = p.lexer.NextToken()
}

func (p *Parser) parseValue() bool {
	switch p.token.type_ {
	case TOKEN_STRING, TOKEN_NUMBER, TOKEN_TRUE, TOKEN_FALSE, TOKEN_NULL:
		p.advance()
		return true
	case TOKEN_LBRACE:
		return p.parseObject()
	case TOKEN_LBRACKET:
		return p.parseArray()
	default:
		return false
	}
}

func (p *Parser) parseObject() bool {
	if p.token.type_ != TOKEN_LBRACE {
		return false
	}
	p.advance()
	if p.token.type_ == TOKEN_RBRACE { // {}
		p.advance()
		return true
	}
	for {
		if p.token.type_ != TOKEN_STRING { // string key
			return false
		}
		p.advance()
		if p.token.type_ != TOKEN_COLON { // followed by colon
			return false
		}
		p.advance()
		if !p.parseValue() { // recursive parse value
			return false
		}
		if p.token.type_ == TOKEN_RBRACE { // terminate
			p.advance()
			return true
		}
		if p.token.type_ != TOKEN_COMMA { // next key follows a comma
			return false
		}
		p.advance()
	}
}

func (p *Parser) parseArray() bool {
	if p.token.type_ != TOKEN_LBRACKET {
		return false
	}
	p.advance()
	if p.token.type_ == TOKEN_RBRACKET {
		p.advance()
		return true
	}
	for {
		if !p.parseValue() {
			return false
		}
		if p.token.type_ == TOKEN_RBRACKET {
			p.advance()
			return true
		}
		if p.token.type_ != TOKEN_COMMA {
			return false
		}
		p.advance()
	}
}
