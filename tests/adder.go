package tests

// START OMIT
func AddSlice(vals []int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

// END OMIT
