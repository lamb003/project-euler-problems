package problem_1

import (
	"github.com/lamb003/project-euler-problems/utils"
	"testing"
)


func TestGenerateNumbers(t*testing.T) {
	numbers := []int{3,5,6,9}
	generatedNumbers := generateMultiplesTill([]int{3,5}, 10)

	if utils.IntSliceIsEqual(numbers, generatedNumbers) != true {

		t.Errorf("Expected: %d got %d", numbers, generatedNumbers)
	}
}

