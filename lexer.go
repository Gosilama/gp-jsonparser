package main

import "slices"

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

	return "", str
}

func lexNumber(str string) (a, b string) {
	return "", str
}

func lexBoolean(str string) (a, b string) {
	return "", str
}

func lexNull(str string) (a, b string) {
	return "", str
}
