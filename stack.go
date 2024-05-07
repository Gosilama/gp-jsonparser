package main

import "fmt"

type StackInt struct {
	s []string
}

// isEmpty() function
func (s *StackInt) IsEmpty() bool {
	length := len(s.s)
	return length == 0
}

// length() function
func (s *StackInt) Length() int {
	length := len(s.s)
	return length
}

// Print() function
func (s *StackInt) Print() {
	length := len(s.s)
	for i := 0; i < length; i++ {
		fmt.Print(s.s[i], " ")
	}
	fmt.Println()
}

// Push() function
func (s *StackInt) Push(value string) {
	s.s = append(s.s, value)
}

// Pop() function
func (s *StackInt) Pop() string {
	length := len(s.s)
	res := s.s[length-1]
	s.s = s.s[:length-1]
	return res
}

// Top() function
func (s *StackInt) Top() string {
	length := len(s.s)
	res := s.s[length-1]
	return res
}
