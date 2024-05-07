package main

import (
	"fmt"
	"log"
	"os"
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

	str := string(file[:])

	fmt.Println(checkParenthesis(str))
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
