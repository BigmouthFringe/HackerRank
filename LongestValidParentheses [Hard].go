package main

import "fmt"

// entry function will calculate result from subLengths
// subLength function will recursively calculate subLength (e.g. '(())', '()', '(((())')

func longestValidParentheses(s string) int {
	var result = 0
	var length = 0

	var n = len(s)
	for i := 0; i < n; i++ {
		for ; (i < n) && (s[i] == '('); {
			var subLength, subShift = subLength(s, i, n)

			var hasPair = (i + subShift < n) && (s[i + subShift] == ')')
			if (subLength != subShift) && !hasPair {
				length = max(length, subLength)
			} else {
				length += subLength
			}

			i += subShift
		}

		result = max(result, length)
		length = 0
	}

	return result
}

func subLength(s string, i, n int) (int, int) {
	var length = 0
	var shift = i

	// TODO: Check if redundant
	if i >= n - 1 {
		return length, 1
	}

	for ; (i < n - 1) && (s[i] == '('); i++ {
		if s[i + 1] != ')' {
			var subLength, subShift = subLength(s, i + 1, n)
			length += subLength
			i += subShift
		} else {
			length += 2; i += 2
			break
		}
	}

	var hasPair = (i - length > 0 && i < n) && (s[i - length] == '(' && s[i] == ')')
	if hasPair {
		length += 2; i++
	}

	shift = i - shift
	return length, shift
}

// Helpers

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(longestValidParentheses("((()") == 2)
	fmt.Println(longestValidParentheses(")()())") == 4)
	fmt.Println(longestValidParentheses("())(())()()))") == 8)
	fmt.Println(longestValidParentheses("(") == 0)
	fmt.Println(longestValidParentheses(")(") == 0)
	fmt.Println(longestValidParentheses("()(())") == 6)
	fmt.Println(longestValidParentheses("()(()") == 2)
	fmt.Println(longestValidParentheses("(()())") == 6)
	fmt.Println(longestValidParentheses("(()(()())()())") == 14)
}
