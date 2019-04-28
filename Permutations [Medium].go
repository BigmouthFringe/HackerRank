package main

import (
	"fmt"
)

var testCases = [][]int{
	{1, 2, 3},
	{1, 2, 3, 4},
	{4, 3, 2, 1, 0},
	{1, 1, 1, 1, 1},
	{5, 7, 8, 9, 222, 123, 556, 45},
	{},
}

func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	perms := make([][]int, 0)
	// TODO: Try to parallelize
	for range nums {
		subPerms := permute(remove(nums, 0))
		for _, perm := range subPerms {
			a := append([]int{nums[0]}, perm...)
			perms = append(perms, a)
		}
		nums = rotate(nums)
	}
	return perms
}

// The functions below do not modify passed slices
// Rotates an array to the left by one element
func rotate(a []int) []int {
	if len(a) == 0 {
		return []int{}
	}
	if len(a) == 1 {
		return a
	}
	res := append([]int{}, a...)
	return append([]int{res[1]}, append(res[2:len(a)], res[0])...)
}
func remove(a []int, i int) []int {
	if len(a) <= 1 {
		return a
	}
	res := append([]int{}, a...)
	return append(res[:i], res[i+1:]...)
}

func main() {
	for _, arr := range testCases {
		perms := permute(arr)
		for _, p := range perms {
			fmt.Println(p)
		}
		fmt.Println()
	}
}
