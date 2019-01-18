package problem_1

import (
	"fmt"
	"github.com/lamb003/project-euler-problems/utils"
)

/*
Problem statement:

	If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
	The sum of these multiples is 23.

	Find the sum of all the multiples of 3 or 5 below 1000.

*/

func generateMultiplesTill(multiples []int, upperLimit int) []int {
	result := make([]int, 0)

	for i := 0; i < upperLimit; i++ {
		for _, multiple := range multiples {
			value := multiple * i

			if utils.IntInSlice(value, result) != true && value < upperLimit {
				result = append(result, value)
			}
		}

	}

	return result
}

func Run() {
	fmt.Println("Problem 1: upper limit 1000")
	result := utils.SumSlice(generateMultiplesTill([]int{3, 5}, 1000))
	fmt.Printf("Answer: %d \n", result)
}
