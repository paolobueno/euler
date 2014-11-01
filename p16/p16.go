/*
2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?
*/

package main

import "math/big"
import "fmt"
import "strconv"

func main() {
	theNumber := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(1000), nil)
	fmt.Println(theNumber)
	arr := []byte(theNumber.String())
	result := 0
	for _, v := range arr {
		char := string(v)
		number, _ := strconv.Atoi(char)
		result += number
	}
	fmt.Println(result)
}