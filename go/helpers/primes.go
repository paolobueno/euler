package helpers

func IsPrime(n int) bool {
	if n <= 3 {
		return n > 1
	}
	if n % 2 == 0 || n % 3 == 0 {
		return false
	}

	for i := 5; i*i <= n; i += 6 {
		if n % i == 0 || n % (i+2)== 0 {
			return false
		}
	}
	return true
}


/*
Returns a list of primes using the sieve of Eratosthenes up to n
 */
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