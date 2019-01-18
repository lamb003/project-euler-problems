package utils

import "testing"


func TestIntSliceIsEqualTrue(t* testing.T) {
	a := []int{1,2,3,4,5}
	b := []int{1,2,3,4,5}

	result := IntSliceIsEqual(a, b)
	if  result != true {
		t.Error("Expected equality got", result)
	}

}

func TestIntSliceIsEqualFalse(t* testing.T) {
	a := []int{1,2,3,4,5}
	b := []int{1,2,3,4,6}

	result := IntSliceIsEqual(a, b)
	if  result != false {
		t.Error("Expected equality got", result)
	}

	b = []int{1,2,3,4,5,6,7}
	result = IntSliceIsEqual(a, b)
	if  result != false {
		t.Error("Expected equality got", result)
	}
}

func TestIntInSliceSuccess(t *testing.T) {
	needle := 5
	haystack := []int {3,4,2,8,10,5,34,1}

	result := IntInSlice(needle, haystack)
	if result != true {
		t.Error("Expected to have found the int in the slice")
	}
}

func TestIntInSliceFailure(t *testing.T) {
	needle := 100
	haystack := []int {3,4,2,8,10,5,34,1}

	result := IntInSlice(needle, haystack)
	if result == true {
		t.Error("Expected to not have found the int in the slice")
	}
}


func TestSumSlice(t*testing.T)  {
	numbers := []int{1,2,3,3,4}

	result := SumSlice(numbers)
	if result != 13 {
		t.Error("Expected 13, got", result)
	}
}
