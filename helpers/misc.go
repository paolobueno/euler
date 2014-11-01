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