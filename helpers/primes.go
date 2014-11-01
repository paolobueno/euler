package helpers

import "math"

func IsPrime(n float64) bool {
	if n <= 3 {
		return n > 1
	}
	if math.Mod(n, 2) == 0 || math.Mod(n, 3) == 0 {
		return false
	}

	for i := float64(5); i*i <= n; i += 6 {
		if math.Mod(n, i) == 0 || math.Mod(n, i+2) == 0 {
			return false
		}
	}
	return true
}
