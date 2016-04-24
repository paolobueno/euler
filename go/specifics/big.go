package specifics

import "math/big"
import "strconv"

func AddDigits(z *big.Int) int{
	arr := []byte(z.String())
	result := 0
	for _, v := range arr {
		char := string(v)
		number, _ := strconv.Atoi(char)
		result += number
	}

	return result
}