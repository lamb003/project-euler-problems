package problem_6

import (
	"fmt"
	"math"
)

/*

	The sum of the squares of the first ten natural numbers is,

	1^2 + 2^2 + ... + 10^2 = 385
	The square of the sum of the first ten natural numbers is,

	(1 + 2 + ... + 10)^2 = 55^2 = 3025
	Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

	Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

*/

func sumOfSquares(number float64) int {
	result := 0.0
	for i := 1.0; i <= number; i++ {
		result = result + math.Pow(i, 2.0)
	}

	return int(result)
}

func squareOfSum(number int) int {
	result := 0
	for i := 1; i <= number; i++ {
		result = result + i
	}
	return int(math.Pow(float64(result), 2.0))
}

func Run() {
	fmt.Println("Problem 6")
	fmt.Printf("Answer: %d\n", squareOfSum(100)-sumOfSquares(100.0))
}
