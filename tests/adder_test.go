package tests

// START OMIT
import "testing"

func TestAddSlice(t *testing.T) {
	a := []int{1, 2, 3, 4, 6, 10}
	if AddSlice(a) != 26 {
		t.Errorf("The sum of %v should be 26", a)
	}
}

// END OMIT
