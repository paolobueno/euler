package helpers

import "math"
/*
Returns channel with divisors for n, fails for small numbers
*/
func Divisors(n int) <-chan int {
	out := make(chan int)
	go func() {
		out <- 1
		limit := int(math.Ceil(math.Sqrt(float64(n)))) // test up to ceil(sqrt(n))
		for i := 2; i <= limit; i++ {
			if n%i == 0 {
				out <- i
				out <- n/i
			}
		}
		close(out)
	}()

	return out
}

func DivisorsSum(n int) int{
	return SumChannel(Divisors(n))
}