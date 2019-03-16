package main

import (
	"fmt"
)

//最长不重复字符串
func lengthOfLongestSubstring(s string) int {
	var first = make(map[uint8]int)
	var second = make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		index, ok := second[s[i]]

		if !ok {
			second[s[i]] = i
		} else {

			if len(first) < len(second) {
				first = make(map[uint8]int)

				for zimu, suoyin := range second {
					first[zimu] = suoyin
				}
			}

			for zimu, suoyin := range second {
				if suoyin <= index {
					delete(second, zimu)
				}
			}
			second[s[i]] = i
		}
	}

	if len(first) > len(second) {
		return len(first)
	} else {
		return len(second)
	}
}

func main() {

	fmt.Println(lengthOfLongestSubstring("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"))
}
