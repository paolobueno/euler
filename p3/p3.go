/*
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
*/
package main

import "github.com/paolobueno/euler/helpers"
import "math"
import "fmt"

func main() {
	n := float64(600851475143)
	for i := math.Ceil(math.Sqrt(n)); i > 5; i-- {
		if helpers.IsPrime(i) && math.Mod(n, i) == 0 {
			fmt.Printf("Found: %f\n", i)
			break
		}
	}
}