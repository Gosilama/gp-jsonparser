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

	fmt.Print(string(file[:]))
}

// func checkParenthesis(s string) bool {
// 	p := map[string]string {
// 		"{": "}",
// 	}

// 	var stack StackInt

// 	if p[s] != nil {
// 		stack.Push(s)
// 	} else {
// 		// check if it's a closed parenthesis and then pop
// 	}
// }
