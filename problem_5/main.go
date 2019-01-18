package problem_5

import "fmt"

/*
	2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

	What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

func bruteForce(number int) int {

	testNumber := number
	found := false
	for found != true {
		testNumber++
		for j := number; j > 0; j-- {
			if testNumber%j != 0 {
				break
			}
			if j == 1 {
				found = true
			}
		}
	}
	return testNumber
}

func Run() {

	fmt.Println("Problem 5")
	fmt.Printf("Answer: %d\n", bruteForce(20))
}
