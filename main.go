package main

import (
	"fmt"
	"log"
	"os"
	// "strings"
)

const JsonQuote = "\""
const JsonComma = ","
const JsonColon = ":"
const JsonLeftBracket = "["
const JsonRightBracket = "]"
const JsonLeftBrace = "{"
const JsonRightBrace = "}"

var JsonWhitespace = []string{" ", "\t", "\b", "\n", "\r"}
var JsonSyntax = []string{JsonComma, JsonColon, JsonLeftBracket, JsonRightBracket, JsonLeftBrace, JsonRightBrace}

const falseLen = len("false")
const trueLen = len("true")
const nullLen = len("null")

func main() {
	if len(os.Args) == 1 {
		log.Fatal("arguments missing")
	}

	args := os.Args[1:]

	// handle reading from stdin later
	file, err := os.ReadFile(args[1])

	if err != nil {
		log.Fatal("could not read file")
	}

	if len(file) <= 0 {
		fmt.Printf("%v is an invalid json file\n", args[1])
		os.Exit(1)
	}

	// str := strings.TrimSpace(string(file[:]))

}

func fromString(str string) {
	fmt.Println(str)
}

// Note: using methods described here : https://notes.eatonphil.com/writing-a-simple-json-parser.html
