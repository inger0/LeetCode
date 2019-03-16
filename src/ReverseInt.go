package main

import (
	"fmt"
	"math"
)

// 反转int数字
func reverse(x int) int {

	target := 0
	flag := 1

	if x < 0 {
		flag = -1
		x = -x
	}

	for x > 0 {
		target = (x % 10) + (target * 10)
		x /= 10
	}

	if target > math.MaxInt32 {
		return 0
	}

	target *= flag

	return target
}

func main() {
	fmt.Print(reverse(-101))
}
