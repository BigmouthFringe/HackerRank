package main

import "fmt"

// Valid parentheses pair:
// Begins with '('
// Ends with ')'
// Contains nothings between border signs

// How to find longest valid parentheses:
// Search for '(' sign
// Check if next sign is ')'
// Check if next sign is '(' and so on until condition is false
// Store quantity of valid pairs in variable
// Repeat and compare new result with the previous one until array ends

func longestValidParentheses(s string) int {
	var result = 0

	var n = len(s)
	var length = 0
	for i := 0; i < n - 1; {
		if s[i] == '(' && s[i + 1] == ')' {
			length += 2
			i += 2
			continue
		} else if length > result {
			result = length
		}

		length = 0
		i++
	}

	if length > result {
		result = length
	}

	return result
}

func main() {
	fmt.Println(longestValidParentheses("(()"))
	fmt.Println(longestValidParentheses(")()())"))
	fmt.Println(longestValidParentheses("())(())()()))"))
	fmt.Println(longestValidParentheses("("))
	fmt.Println(longestValidParentheses(")("))
	fmt.Println(longestValidParentheses("()(())")) // Expected: 6
}
