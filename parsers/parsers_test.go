package parsers

import "testing"

func sliceEq(a, b []int) bool {

	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestInferColumnOffsets(t *testing.T) {
	cases := []struct {
		rows    []string
		offsets []int
	}{
		{[]string{"aaa b   cc", "1   234 5 "}, []int{0, 4, 8}},
		{[]string{"Wicket, Steve   Mendelssohnstraat 54d 3423 ba  0313-398475           93400 19640603",
			"Smith, John     Børkestraße 32        87823    +44 728 889838       989830 19990920",
			"Friendly, User  Sint Jansstraat 32    4220 EE  0885-291029            6360 19800810",
			"Name            Address               Postcode Phone          Credit Limit Birthday"},
			[]int{0, 16, 38, 47, 62, 75}},
	}

	for _, c := range cases {
		offsets := inferOffsets(c.rows)
		if !sliceEq(offsets, c.offsets) {
			t.Errorf("inferOffsets(%q) == %v, want %v", c.rows, offsets, c.offsets)
		}
	}
}
