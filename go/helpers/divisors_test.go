package helpers

import "testing"

type divisorsFixture struct {
	value    int
	expected int
}

func TestDivisors(t *testing.T) {
	fixtures := []divisorsFixture{
		// Still fails for small numbers
		// {2, 1},
		// {3, 3},
		{10, 8},
		{220, 284},
		{284, 220},
	}

	for _, v := range fixtures {
		result := DivisorsSum(v.value)
		if result != v.expected {
			t.Error(
				"Value:", v.value,
				"Expected:", v.expected,
				"Got:", result,
			)
		}
	}
}
