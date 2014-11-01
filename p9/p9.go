/*
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/
package main

import "fmt"
import "math"

func main() {
	for a := 1; a < 1000; a++ {
		for b := a; b < 1000; b++ {
			c := math.Sqrt(float64((a*a)+(b*b)))
			if c == math.Floor(c) && int(c) > b {
				if a + b + int(c) == 1000 {
					fmt.Println(
						"a:", a,
						"b:", b,
						"c:", c,
					)
					fmt.Println(a * b * int(c))
					return
				}
			}
			
		}
	}
}