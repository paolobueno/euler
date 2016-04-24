/*
Surprisingly there are only three numbers that can be written as the sum of fourth powers of their digits:

1634 = 1^4 + 6^4 + 3^4 + 4^4 
8208 = 8^4 + 2^4 + 0^4 + 8^4 
9474 = 9^4 + 4^4 + 7^4 + 4^4 
As 1 = 14 is not a sum it is not included.

The sum of these numbers is 1634 + 8208 + 9474 = 19316.

Find the sum of all the numbers that can be written as the sum of fifth powers of their digits.
*/
package main

import "fmt"
import "github.com/paolobueno/euler/helpers"

func pow5(n int) int {
	return n*n*n*n*n
}

func main() {
	result := 0
	for i := 2; i < 1e6; i++ {
		sum := 0
		for n := range helpers.EachDigit(i) {
			sum += pow5(n)
		}
		if sum == i {
			result += i
			fmt.Println(i)
		}
	}

	fmt.Println("Result: ", result)
}