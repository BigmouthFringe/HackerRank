package main

import "fmt"

// this approach will use this stack
type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}
)

func NewStack() *Stack {
	return &Stack{nil,0}
}

func (this *Stack) Pop() interface{} {
	if this.length == 0 {
		return nil
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value
}

func (this *Stack) Push(value interface{}) {
	n := &node{value,this.top}
	this.top = n
	this.length++
}

func (this *Stack) Empty() bool {
	return this.length == 0
}

// if element of an array is '(' - push it
// if stack is not empty and element is ')' - pop, length + 2
func longestValidParentheses(s string) int {
	var result = 0
	var length = 0
	var stack = NewStack()

	var n = len(s)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack.Push(s[i])
		} else if !stack.Empty() && s[i] == ')' {
			stack.Pop()
			length += 2

			if stack.Empty() {
				result = max(result, length)
			}
		} else {
			result = max(result, length)
			length = 0
		}
	}

	if result == 0 {
		result = length
	}

	return result
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
	fmt.Println(longestValidParentheses("(())(()") == 4)
	fmt.Println(longestValidParentheses("(()())") == 6)
	fmt.Println(longestValidParentheses("(()(()())()())") == 14)
	fmt.Println(longestValidParentheses("(()(((()") == 2)
}
