package helpers

func Fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		ret := x
		x, y = y, x+y
		return ret
	}
}

func Eratosthenes(n int) []int {
	arr := make([]bool, n + 1)
	// initialize to true
	for i := range arr {
		arr[i] = true
	}

	primeCount := 0
	for i := 2; i < len(arr); i++ {
		if arr[i] { // is still prime
			primeCount++
			for mult := i + i; mult <= n; mult += i {
				arr[mult] = false
			}
		}
	}

	result := make([]int, primeCount)
	result[0] = 2 // 2 for first prime
	resultCursor := 1
	for i := 3; i < len(arr); i++ {
		isPrime := arr[i]
		if isPrime {
			result[resultCursor] = i
			resultCursor++
		}
	}

	return result
}
