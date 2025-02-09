package main

import (
	"slices"
	"strconv"
	"strings"
)

// type Number interface {
// 	int64 | float64
// }

func lexer(str string) {
	var tokens []string

	for len(str) > 0 {
		jsonStr, s := lexString(str)

		if jsonStr != "" {
			tokens = append(tokens, jsonStr)
			continue
		}

		if slices.Contains(JsonWhitespace, string(s[0])) {
			s = s[1:]
		} else if slices.Contains(JsonSyntax, string(s[0])) {
			tokens = append(tokens, string(s[0]))
		} else {
			panic("Unexpected character: {}")
		}
	}
}

func lexString(str string) (a, b string) {
	jsonStr := ""

	if string(str[0]) == JsonQuote {
		str = str[1:]
	} else {
		return "", str
	}

	for _, val := range str {
		if string(val) == JsonQuote {
			return jsonStr, str[len(jsonStr)+1:]
		} else {
			jsonStr += string(val)
		}
	}

	panic("Expected end-of-string quote")
}

// I'll update this to use the interface constraint approach when I understand generics better
func lexNumber(str string) (n any, b string) {
	jsonNumber := ""

	numberChars := []string{"-", "e", "."}

	for val := range 10 {
		numberChars = append(numberChars, string(val))
	}

	for _, val := range str {
		if slices.Contains(numberChars, string(val)) {
			jsonNumber += string(val)
		} else {
			break
		}
	}

	leftover := str[len(jsonNumber):]

	if len(jsonNumber) == 0 {
		return nil, str
	}

	if strings.Contains(jsonNumber, ".") {
		val, err := strconv.ParseFloat(jsonNumber, 64)

		if err != nil {
			panic("invalid float conversion")
		}

		return val, leftover
	}

	val, _ := strconv.ParseInt(jsonNumber, 10, 64)
	return val, leftover
}

// this should return a boolean and a string
func lexBoolean(str string) (a any, b string) {
	strLen := len(str)

	if strLen >= trueLen && str[:trueLen] == "true" {
		return true, str[trueLen:]
	}

	if strLen >= falseLen && str[:falseLen] == "false" {
		return false, str[falseLen:]
	}

	return nil, str
}

func lexNull(str string) (a any, b string) {
	strLen := len(str)

	if strLen >= nullLen && str[:nullLen] == "null" {
		return true, str[nullLen:]
	}

	return nil, str
}
