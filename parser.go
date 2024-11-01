package main

import "strconv"

func Parser(lexToken []string) ([]string, bool) {
	var result bool = false
	if len(lexToken) < 1 {
	  return lexToken, result
	}
	token := lexToken[0]
	if token == LeftCurlyBrace {
	  lexToken = lexToken[1:]
	  lexToken, result = IsValidObject(lexToken)
	} else if token == LeftSquareBrace {
	  lexToken = lexToken[1:]
	  lexToken, result = IsValidArray(lexToken)
	} else {
	  result = IsValidString(token) || IsValidNumber(token) || IsValidBoolean(token) || IsValidNull(token)
	  lexToken = lexToken[1:]
	}
	  return lexToken, result
  }

// function to validate weather `lexToken` contains a valid object
func IsValidObject(lexToken []string) ([]string, bool) {
	if len(lexToken) == 0 {
		return lexToken, false
	}
	if lexToken[0] == RightCurlyBrace {
	return lexToken[1:], true
	}
	for {
		// for json key
		if len(lexToken) == 0 {
			return lexToken, false
		}
		json_key := lexToken[0]
		lexToken = lexToken[1:]
		if !IsValidString(json_key) {
			return lexToken, false
		}

		// for json key value colon
		if len(lexToken) == 0 {
			return lexToken, false
		}
		json_colon := lexToken[0]
		lexToken = lexToken[1:]
		if json_colon != Colon {
			return lexToken, false
		}

		// for json value
		if len(lexToken) == 0 {
			return lexToken, false
		}
		json_value := lexToken[0]
		lexToken = lexToken[1:]
		if json_value == LeftCurlyBrace {
			var tempResult bool
			lexToken, tempResult = IsValidObject(lexToken)
			if !tempResult {
				return lexToken, false
			}
		} else if json_value == LeftSquareBrace {
			var tempResult bool
			lexToken, tempResult = IsValidArray(lexToken)
			if !tempResult {
				return lexToken, false
			}
		} else if !IsValidString(json_value) && !IsValidNumber(json_value) && !IsValidBoolean(json_value) && !IsValidNull(json_value) {
			return lexToken, false
		}

		if len(lexToken) == 0 {
			return lexToken, false
		}
		json_comma_or_curly_brace := lexToken[0]
		lexToken = lexToken[1:]
		if json_comma_or_curly_brace == RightCurlyBrace {
			break
		}
		if json_comma_or_curly_brace != Comma {
			return lexToken, false
		}
	}
	return lexToken, true
}

// checks weather `lexToken` contains valid array
func IsValidArray(lexToken []string) ([]string, bool) {
if len(lexToken) == 0 {
	return lexToken, false
}
token := lexToken[0]
if token == RightSquareBrace {
	return lexToken[1:], true
}

for {
	var tempResult bool
	lexToken, tempResult = Parser(lexToken)
	if !tempResult {
	return lexToken, false
	}
	token := lexToken[0]
	if token == RightSquareBrace {
	return lexToken[1:], true
	} else if token != Comma {
	return lexToken[1:], false
	} 
	lexToken = lexToken[1:]
}

return lexToken, true
}

// weather the input is a valid json stirng type
func IsValidString(input string) bool {
	inputLen := len(input)
if inputLen < 2 {
	return false
}
	var prevChar rune
	if input[0] != DoubleQuoteByte || input[inputLen-1] != DoubleQuoteByte {
		return false
	}
	for _, ch := range input {
		if prevChar == BackSlashRune {
			var validAfterBackSlash bool = false
			for _, allowedChar := range AllowedCharsAfterEscapeChar {
				if allowedChar == ch {
					validAfterBackSlash = true
					break
				}
			}
			if !validAfterBackSlash {
				return false
			}
			prevChar = 0
		} else {
			prevChar = ch
		}
	}
	return true
}

// checks weather the input string can be converted to a valid number or not
func IsValidNumber(input string) bool {
	_, atoiErr := strconv.Atoi(input)
	_, floatErr := strconv.ParseFloat(input, 64)
	if atoiErr != nil && floatErr != nil {
		return false
	}
	return true
}

// weather the input is a valid json bool data type
func IsValidBoolean(input string) bool {
	return input == "true" || input == "false"
}

// weather input is a valid null
func IsValidNull(input string) bool {
	return input == "null"
}