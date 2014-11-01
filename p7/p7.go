/*
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/
package main

import (
	"github.com/paolobueno/euler/helpers"
	"fmt"
)

func main() {
	n, count := float64(0), 0
	for count < 10001 {
		if helpers.IsPrime(n) {
			count++
		}
		n++
	}
	fmt.Printf("Prime found: %d, #%d\n", n, count)
}
