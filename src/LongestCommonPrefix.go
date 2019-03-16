package main

import "fmt"

// 最长公共前缀
func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}

	target := ""
	index := 0
	for true {
		continueFlag := true
		if index >= len(strs[0]) {
			break
		}

		for i := 1; i < length; i++ {
			subStr := strs[i]
			if index >= len(subStr) {
				continueFlag = false
				break
			}

			if strs[0][index] != subStr[index] {
				continueFlag = false
				break
			}
		}

		if !continueFlag {
			break
		}

		target += fmt.Sprintf("%c", strs[0][index])

		index++
	}

	return target
}

func main() {

	a := []string{"ddf"}

	fmt.Println(longestCommonPrefix(a))
}
