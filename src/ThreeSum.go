package main

import (
	"fmt"
	"sort"
)

//给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
//注意：答案中不可以包含重复的三元组。
func threeSum(nums []int) [][]int {

	var target [][]int
	sort.Ints(nums)
	length := len(nums)

	if length < 3 {
		return target
	}

	for i := 0; i < length-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		head := i + 1
		tail := length - 1

		tmpTarget := -nums[i]
		for head < tail {
			if (nums[head] + nums[tail]) == tmpTarget {

				target = append(target, []int{nums[head], nums[i], nums[tail]})
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

	return target
}

func main() {
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}
