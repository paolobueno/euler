/*
We shall say that an n-digit number is pandigital if it makes use of all the
digits 1 to n exactly once; for example, the 5-digit number, 15234, is 1 through 5 pandigital.

The product 7254 is unusual, as the identity, 39 × 186 = 7254, containing
multiplicand, multiplier, and product is 1 through 9 pandigital.

Find the sum of all products whose multiplicand/multiplier/product identity can
be written as a 1 through 9 pandigital.
*/
package main

import "fmt"
import "strconv"

func main() {
	// Only len(1) * len(4) and len(2) * len(3) numbers can result in
	// len(4) numbers
	for i := 0; i < 10; i++ {
		for j := i; j < 10000; j++ {
			if check(i,j) {
				fmt.Println(i,j, i*j)
			}
		}
	}
	for i := 0; i < 100; i++ {
		for j := i; j < 1000; j++ {
			if check(i,j) {
				fmt.Println(i,j, i*j)
			}
		}
	}
}

func check(i, j int) bool {
	mult := i * j
	return mult < 10000 && isPandigital(strconv.Itoa(i) + strconv.Itoa(j) + strconv.Itoa(mult))
}

func isPandigital(str string) bool {
	hasDigits := uint16(0)
	for _, rn := range str {
		if rn == '0' {
			return false
		}
		number := uint(rn - '1')
		if hasDigits & (1 << number) > 0 {
			return false // repeated number
		}
		hasDigits |= 1 << number
	}
	return hasDigits == 0x1ff // 8 bits [1..9]
}