package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {

	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	if dividend == 0 {
		return 0
	}

	sign := 1
	if dividend < 0 {
		sign *= -1
	}

	if divisor < 0 {
		sign *= -1
	}

	if dividend == 0 {
		return 0
	}

	tmpA, tmpB := abs64(dividend), abs64(divisor)
	var ret int
	for tmpA >= tmpB {
		tmp, m := tmpB, 1
		for tmpA >= tmp<<1 {
			tmp, m = tmp<<1, m<<1
		}
		tmpA -= tmp
		ret += m
	}

	//dvd, dvs := abs64(dividend), abs64(divisor)
	//var ret int
	//for dvd >= dvs {
	//	tmp, m := dvs, 1
	//	for dvd >= tmp<<1 {
	//		tmp, m = tmp<<1, m<<1
	//	}
	//	dvd -= tmp
	//	ret += m
	//}

	return ret * sign
}

func abs64(n int) int64 {
	ret := int64(n)
	if ret < 0 {
		return -ret
	}
	return ret
}

func main() {

	m := 111
	fmt.Println(m << 1)

	fmt.Println(divide(10, 3))
}
