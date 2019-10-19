package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	reach := reachingPoints(1, 1, 3, 5)
	fmt.Println(reach, time.Since(start))

	start = time.Now()
	reach = reachingPoints(1, 1, 2, 2)
	fmt.Println(reach, time.Since(start))

	start = time.Now()
	reach = reachingPoints(1, 1, 1, 1)
	fmt.Println(reach, time.Since(start))
}

func reachingPoints(sx int, sy int, tx int, ty int) bool {
	if sx == tx && sy == ty { return true }
	var stack [][2]int
	stack = append(stack, [2]int{sx+sy, sy}, [2]int{sx, sx+sy})
	return traverse(&stack, tx, ty)
}

func traverse(stack *[][2]int, tx int, ty int) bool {
	for len(*stack) != 0 {
		tuple := (*stack)[len(*stack)-1]
		*stack = (*stack)[:len(*stack)-1]
		sx, sy := tuple[0], tuple[1]
		if sx == tx && sy == ty { return true }
		if !(sx > tx || sy > ty) {
			*stack = append(*stack, [2]int{sx+sy, sy}, [2]int{sx, sx+sy})
		}
	}
	return false
}

// stack space exceeded
//func reachingPoints(sx int, sy int, tx int, ty int) bool {
//	if sx == tx && sy == ty {
//		return true
//	}
//	if sx > tx || sy > ty {
//		return false
//	}
//
//	left, right := reachingPoints(sx+sy, sy, tx, ty), reachingPoints(sx, sx+sy, tx, ty)
//	if left || right {
//		return true
//	}
//	return false
//}
