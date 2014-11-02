package helpers

const MaxUint = ^uint(0) 
const MinUint = 0 
const MaxInt = int(MaxUint >> 1) 
const MinInt = -MaxInt - 1

func Max(one, other int) int {
	if one > other {
		return one
	} else {
		return other
	}
}

func Min(one, other int) int {
	if one < other {
		return one
	} else {
		return other
	}
}

func MaxInSlice(slice []int) int{
	max := MinInt
	for _, v := range slice {
		if max < Max(max, v) {
			max = v
		}
	}
	return max
}

func ChanEqSlice(c <-chan int, slice []int) bool {
	count := 0
	for i := range c{
		if count >= len(slice) {
			return false
		}

		if i != slice[count] {
			return false
		}
		count++
	}
	return true
}

func EachDigit(n int) <-chan int {
	c := make(chan int)
	go func() {
		for i := n; i > 0; i = i / 10 {
			c <- i % 10
		}
		close(c)
	}()
	return c
}