package tests

// START OMIT
import "testing"

type addTest struct {
	in  []int
	out int
}

var data = []addTest{
	addTest{in: []int{1, 2, 3}, out: 6},
	addTest{in: []int{1, 2, 3}, out: 5},
	addTest{in: []int{1, 2, 3}, out: 4},
}

func TestAddSlice(t *testing.T) {
	for _, d := range data {
		actual := AddSlice(d.in)
		if actual != d.out {
			t.Errorf("Expected %d, got %v", d.out, actual) // HLerror
		}
	}
}

// END OMIT
