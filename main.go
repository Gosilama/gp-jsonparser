package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

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

	str := strings.TrimSpace(string(file[:]))
	bStr, kvStr := getParenthesis(str)

	valid := checkParenthesis(bStr)

	if !valid {
		fmt.Printf("%v is an invalid json file\n", args[1])
		os.Exit(1)
	}

	// kvStrArr := strings.Split(kvStr, ":")

	// for _, val := range kvStrArr {
	// 	if reflect.TypeOf(val) != "string" {

	// 	}
	// }
}

func getParenthesis(str string) (string, string) {
	// str = strings.TrimSpace(str)
	strLen := len(str)

	return string(str[0]) + "" + string(str[strLen-1]), str[1 : strLen-1]
}

func checkParenthesis(str string) bool {
	// can be collapsed into a nested map
	op := map[string]bool{
		"{": true,
		"[": true,
		"(": true,
	}

	cp := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}

	var stack Stack

	for _, c := range str {
		s := string(c)
		if op[s] {
			stack.Push(s)
		} else if stack.IsEmpty() {
			return false
		} else {
			// check if it's a closed parenthesis and then pop
			val := stack.Pop()

			if cp[s] != val {
				return false
			}
		}
	}

	return true
}
