package main

import "fmt"

// 盛水最多的木桶
func maxArea(height []int) int {

	//length := len(height)
	//var result = 0
	//for i := 0; i < length-1; i++ {
	//
	//	for j := i + 1; j < length; j++ {
	//
	//		tmp := height[j] * (j - i)
	//
	//		if height[i] < height[j] {
	//			tmp = height[i] * (j - i)
	//		}
	//
	//		if result < tmp {
	//			result = tmp
	//		}
	//	}
	//}

	length := len(height)
	var result = 0
	var head = 0
	var tail = length - 1
	for head < tail {
		tmp := 0

		if height[head] > height[tail] {
			tmp = height[tail] * (tail - head)
			tail--

		} else {
			tmp = height[head] * (tail - head)
			head++
		}

		if result < tmp {
			result = tmp
		}
	}

	return result
}

func main() {
	fmt.Print(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
