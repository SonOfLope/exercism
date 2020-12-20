package diffsquares

import "math"

// SumOfSquares will caculate the sum of the squares of every number until N and return that value
func SumOfSquares(n int) int {
	sumOfSq := 0
	for count := 0; count <= n; count++ {
		sumOfSq = sumOfSq + int(math.Pow(float64(count), 2))
	}
	return sumOfSq
}

// SquareOfSum will caculate the square of the sum of every number to N and return that value
func SquareOfSum(n int) int {
	sqOfSum := 0
	for count := 0; count <= n; count++ {
		sqOfSum = sqOfSum + count
	}
	return sqOfSum * sqOfSum
}

// Difference will calculate the difference between sqOfSum and sumOfSq and return that value
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
