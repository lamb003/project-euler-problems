package problem_4

import "fmt"

/*
	A palindromic number reads the same both ways.
	The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

	Find the largest palindrome made from the product of two 3-digit numbers.

*/

func reverseInt(number int) int {
	result := 0
	lastDigit := 1

	for number > 0 {
		lastDigit = number % 10
		result = result*10 + lastDigit
		number = number / 10
	}
	return result
}

func isPalindrome(value int) bool {
	return value == reverseInt(value)
}

func Run() {
	fmt.Println("Problem 4: ")

	largestPal := 0

	for i := 999; i > 99; i-- {
		for j := i; j > 99; j-- {
			if i*j < largestPal {
				break // It just gets smaller from here for each following value for j
			}

			if isPalindrome(i * j) {
				largestPal = i * j
			}
		}
	}

	fmt.Printf("Answer: %d\n", largestPal)
}
