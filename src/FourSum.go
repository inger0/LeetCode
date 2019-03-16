package main

import (
	"fmt"
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	var result [][]int
	length := len(nums)
	if length < 4 {
		return result
	}

	sort.Ints(nums)
	for i := 0; i < length-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < length-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			head := j + 1
			tail := length - 1

			tmpTarget := target - (nums[i] + nums[j])
			for head < tail {
				if (nums[head] + nums[tail]) == tmpTarget {
					result = append(result, []int{nums[i], nums[j], nums[head], nums[tail]})
					head++
					tail--
					for head < tail && nums[head] == nums[head-1] {
						head++
					}
					for head < tail && nums[tail] == nums[tail+1] {
						tail--
					}
				} else if (nums[head] + nums[tail]) < tmpTarget {
					head++
				} else {
					tail--
				}
			}
		}
	}

	return result
}

func main() {

	fmt.Println(fourSum([]int{0, 1, 1, 1, 1, 1, 2, 3}, 5))
}
