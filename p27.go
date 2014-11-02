/*
Euler discovered the remarkable quadratic formula:

n² + n + 41

It turns out that the formula will produce 40 primes for the consecutive values n = 0 to 39.
However, when n = 40, 402 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41, and
certainly when n = 41, 41² + 41 + 41 is clearly divisible by 41.

The incredible formula  n² − 79n + 1601 was discovered, which produces 80 primes
for the consecutive values n = 0 to 79.
The product of the coefficients, −79 and 1601, is −126479.

Considering quadratics of the form:

n² + an + b, where |a| < 1000 and |b| < 1000

where |n| is the modulus/absolute value of n
e.g. |11| = 11 and |−4| = 4
Find the product of the coefficients, a and b, for the quadratic expression that
produces the maximum number of primes for consecutive values of n, starting with n = 0.
*/
package main

import "fmt"
import "github.com/paolobueno/euler/helpers"

type results struct{
	limit int
	a int
	b int
}

func main() {
	result := results{0, 0, 0}
	for a := -1000; a <= 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			for n := 0; helpers.IsPrime(formula(n,a,b)); n++ {
				if n > result.limit {
					result.limit = n
					result.a, result.b = a, b
				}
			}
		}
	}
	fmt.Println(result)
	fmt.Println(result.a*result.b)
}

func formula(n, a, b int) int {
	return (n*n) + a*n + b
}