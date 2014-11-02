/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/
package main

import "github.com/paolobueno/euler/helpers"
import "fmt"

func main() {
	const limit int = 2e6;
	primes := helpers.Eratosthenes(limit)
	sum := 0
	for _, v := range primes {
		sum += v
	}
	fmt.Println("result:", sum)
}