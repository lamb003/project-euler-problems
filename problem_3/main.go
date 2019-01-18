package problem_3

import "fmt"

/*
	The prime factors of 13195 are 5, 7, 13 and 29.

	What is the largest prime factor of the number 600851475143 ?
*/

func factor(number int) int {

	i := 2
	// traverse factor tree
	for i*i < number {
		if number%i == 0 {
			number = number / i
		}
		i = i + 1
	}
	return number
}

func Run() {
	result := factor(600851475143)

	fmt.Println("Problem 3 ")
	fmt.Printf("Answer: %d\n", result)
}
