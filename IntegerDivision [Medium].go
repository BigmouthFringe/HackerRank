package main

import (
	"fmt"
	"math"
	. "strconv"
)

func divide(dividend int, divisor int) int {
	var resultSign = sign(dividend) * sign(divisor)
	dividend *= sign(dividend)
	divisor *= sign(divisor)

	var digits = sliceInteger(dividend)
	var currentInteger = digits[0]
	var resultDigits []int
	var n = len(digits)
	for i := 0; i < n - 1; {
		for ; (currentInteger < divisor) && (i < n - 1); i++ {
			var integersToJoin = []int{currentInteger, digits[i + 1]}
			currentInteger = joinIntegers(integersToJoin)
		}

		if currentInteger < divisor {
			break
		}

		// NOTE: This loop might be a hot spot
		var subtractor = divisor
		var subtractorTimes = 1
		for currentInteger >= (subtractor + divisor) {
			subtractor += divisor
			subtractorTimes++
		}

		currentInteger = currentInteger - subtractor
		resultDigits = append(resultDigits, subtractorTimes)
		var appendZero = (i < n - 1) && (joinIntegers([]int{currentInteger, digits[i + 1]}) < divisor)
		if appendZero {
			resultDigits = append(resultDigits, 0)
			i++
			currentInteger = digits[i]
		}
	}

	if (len(digits) == 1) && (dividend >= divisor) {
		var subtractor = divisor
		var subtractorTimes = 1
		for currentInteger >= (subtractor + divisor) {
			subtractor += divisor
			subtractorTimes++
		}
		resultDigits = append(resultDigits, subtractorTimes)
	}

	var result = joinIntegers(resultDigits) * resultSign
	if result > math.MaxInt32 || result < math.MinInt32 {
		return math.MaxInt32
	}
	return result
}

// Helpers

func sliceInteger(integer int) []int {
	var intString = Itoa(integer)
	var intRunes = []rune(intString)

	var n = len(intRunes)
	var digits = make([]int, n)
	for i := 0; i < n; i++ {
		digits[i] = int(intRunes[i] - '0')
	}

	return digits
}

func joinIntegers(digits []int) int {
	if digits == nil {
		return 0
	}
	if len(digits) == 1 {
		return digits[0]
	}

	var n = len(digits)
	var integer = digits[0]
	for i := 0; i < n - 1; i++ {
		integer = (integer * 10) + digits[i + 1]
	}

	return integer
}

func sign(x int) int {
	if x < 0 {
		return -1
	} else {
		return 1
	}
}

func main() {
	fmt.Println(divide(-7, 1))
	fmt.Println(divide(1004958205, -2137325331))
	fmt.Println(divide(2147483647, 2147483647))
	fmt.Println(divide(0, 1))
	fmt.Println(divide(-2147483648, 1))
	fmt.Println(divide(-1060849722, 99958928))
	fmt.Println(divide(-1021989372, -82778243))
}
