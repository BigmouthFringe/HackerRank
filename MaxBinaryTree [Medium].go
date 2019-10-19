package main

import (
	"fmt"
	"time"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func main() {
	start := time.Now()
	constructMaximumBinaryTree([]int{3,2,1,6,0,5})
	fmt.Println(time.Since(start))
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	maxInd := 0
	for i := range nums {
		if nums[maxInd] < nums[i] {
			maxInd = i
		}
	}

	var left, right *TreeNode
	if maxInd > 0 {
		left = constructMaximumBinaryTree(nums[0:maxInd])
	}
	if maxInd+1 < len(nums) {
		right = constructMaximumBinaryTree(nums[maxInd+1:])
	}

	return &TreeNode{nums[maxInd], left, right}
}
