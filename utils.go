package main

func IntAbs(val int) int {
	if val >= 0 {
		return val
	}
	return -val
}

func IntMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func IntMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func StartIndex(delta int, first int, second int) int {
	index := first
	if delta > 0 {
		index = IntMin(first, second)
	} else if delta < 0 {
		index = IntMax(first, second)
	}
	return index
}

func Sign(delta int) int {
	sign := 0
	if delta > 0 {
		sign = 1
	} else if delta < 0 {
		sign = -1
	}
	return sign
}
