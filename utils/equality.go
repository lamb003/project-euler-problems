package utils


// Test for equality in 2 int slices
func IntSliceIsEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// Test if an int is in a slice of ints
func IntInSlice(needle int, haystack []int)  bool {
	for _, value := range haystack {
		if value == needle {
			return true
		}
	}
	return false

}

func SumSlice(data []int) int  {
	result := 0
	for _, value := range data{
		result = result + value
	}
	return result
}