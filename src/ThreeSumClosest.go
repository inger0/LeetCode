package main

import (
	"fmt"
	"math"
	"sort"
)

// 求三个数之和最接近于目标
func threeSumClosest(nums []int, target int) int {
	length := len(nums)

	if length < 3 {
		return 0
	}
	sort.Ints(nums)

	result := nums[0] + nums[1] + nums[2]
	for i := 1; i < length-2; i++ {
		left := i + 1
		right := length - 1
		for left < right {
			sum := nums[left] + nums[right] + nums[i]

			if sum == target {
				return sum
			}

			if sum > target {
				right--
			} else {
				left++
			}

			if math.Abs(float64(target-sum)) < math.Abs(float64(target-result)) {
				result = sum
			}
		}
	}

	return result
}

func main() {
	fmt.Println(threeSumClosest([]int{-1, 2, 1, -4}, 1))
}
