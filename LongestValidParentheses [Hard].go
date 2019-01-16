package main

import "fmt"

// Valid parentheses pair:
// Begins with '('
// Ends with ')'
// May contain a valid pair of parentheses

// Valid parentheses substrings:
// (()) = complex substring, ()() = simple substrings

// We assume that every new possible substring is complex

// complexBracesLength should verify brackets and return 0 if they're not valid

func longestValidParentheses(s string) int {
	var result = 0

	var n = len(s)
	var length = 0
	for i := 0; i < n - 1; i++ {
		if s[i] == '(' {
			var subLength = complexBracesLength(s, i, n)
			if subLength > 0 {
				length += subLength
				i += subLength - 1
			} else {
				result = max(result, length)
				length = 0
				continue
			}
		} else {
			continue
		}

		if (i < n - 1) && (s[i + 1] != '(') {
			result = max(result, length)
			length = 0
			i++
		}
	}

	if length > result {
		result = length
	}

	return result
}

func complexBracesLength(s string, i int, n int) int {
	var length = 0
	var count = 0

	// Step 1
	if s[i] == '(' {
		count++
		for ; (i < n - 1) && s[i + 1] != ')'; i++ {
			count++
		}
	} else {
		return 0
	}

	// Step 2
	for j := 0; j < count; j++ {
		if (i < n - 1) && (s[i + 1] != '(') {
			length += 2
			i++
		} else {
			// Parentheses are not valid
			return 0
		}
	}

	return length
}

func entry(s string) int {
	var result = 0
	var length = 0
	var shift = 0

	var n = len(s)
	for i := 0; i < n; i++ {
		for ; i < n && s[i] == '('; shift++ {
			length++
			var subLength, subShift = subLength(s, i, n)
			length += subLength; shift += subShift
			i += subShift

			var hasPair = (i - length > 0 && i < n) && (s[i - length] == '(' && s[i] == ')')
			if !hasPair {
				length--
			}
		}

		result = max(result, length)
		length = 0; shift = 0
	}

	return result
}

func subLength(s string, i int, n int) (int, int) {
	var lenResult = 0
	var shiftResult = i

	var length = 0
	for i < n - 1 {
		if (s[i] == '(') && (s[i + 1] != ')') {
			i++
			var subLength, shift = subLength(s, i, n)
			length += subLength
			i += shift
		} else {
			i++
			lenResult = 2; shiftResult = 2
			return lenResult, shiftResult
		}

		i++
		lenResult = max(lenResult, length)
		length = 0
	}

	shiftResult = i - shiftResult
	return lenResult, shiftResult
}

// Helpers

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(entry("((()") == 2)
	fmt.Println(entry(")()())") == 4)
	fmt.Println(entry("())(())()()))") == 8)
	fmt.Println(longestValidParentheses("(") == 0)
	fmt.Println(longestValidParentheses(")(") == 0)
	fmt.Println(longestValidParentheses("()(())") == 6)
	fmt.Println(longestValidParentheses("()(()") == 2)
	fmt.Println(longestValidParentheses("(()(()())()())") == 14) // Use recursion
}
