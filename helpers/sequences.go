package helpers

import "math/big"

func Fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		ret := x
		x, y = y, x+y
		return ret
	}
}

var fibBigCache map[int]big.Int

func FibonacciBig(i int) big.Int {
	if fibBigCache == nil {
		fibBigCache = make(map[int]big.Int)
		fibBigCache[0] = *big.NewInt(0)
		fibBigCache[1] = *big.NewInt(1)
	}

	_, ok := fibBigCache[i]
	if !ok {
		res := new(big.Int)
		k := new(big.Int)
		j := new(big.Int)
		*k = FibonacciBig(i-1)
		*j = FibonacciBig(i-2)
		res.Add(k,j)
		fibBigCache[i] = *res
	}
	return fibBigCache[i]
}

func SumChannel(c <- chan int) int {
	sum := 0
	for i := range c {
		sum += i
	}
	return sum
}

func LexicographicPermute(arr []int) bool {
	// http://en.wikipedia.org/wiki/Permutation#Algorithms_to_generate_permutations

	// Find the largest index k such that a[k] < a[k + 1]. If no such index exists, the permutation is the last permutation.
	k := len(arr)
	for k = len(arr) - 2; k > -1 && (arr[k] >= arr[k+1]) ; k-- {
	}
	if k == -1 {
		return false
	}
	// Find the largest index l greater than k such that a[k] < a[l].
	l := 0
	for l = len(arr) - 1; l > k  && (arr[k] >= arr[l]); l-- {
	}
	// Swap the value of a[k] with that of a[l].
	Swap(arr, k, l)
	// Reverse the sequence from a[k + 1] up to and including the final element a[n].
	Reverse(arr[k+1:])

	return true
}

func Swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func Reverse(a []int) {
	for i1, i2 := 0, len(a)-1; i1 < i2;  {
		a[i1], a[i2] = a[i2], a[i1]
		i1++
		i2--
	}
}