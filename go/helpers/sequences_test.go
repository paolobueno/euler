package helpers

import "testing"

type eratosthenesFixture struct {
	value int
	expected []int
}

func TestEratosthenes(t *testing.T) {
	fixtures := []eratosthenesFixture{
		{2, []int{2}},
		{5, []int{2,3,5}},
		{10, []int{2,3,5,7}},
		{300, []int{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101,103,107,109,113,127,131,137,139,149,151,157,163,167,173,179,181,191,193,197,199,211,223,227,229,233,239,241,251,257,263,269,271,277,281,283,293}},
	}

	for _, f := range fixtures {
		result := Eratosthenes(f.value)
		if !sliceEqual(result, f.expected) {
			t.Error(
				"Value:", f.value,
				"Expected:", f.expected,
				"Got:", result,
			)
		}
	}
}

func sliceEqual(a, b []int) bool {
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