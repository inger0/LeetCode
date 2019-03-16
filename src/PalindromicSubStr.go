package main

import (
	"fmt"
)

// 最长回文字符串
func longestPalindrome(s string) string {

	var result string
	var palStr string

	var length = len(s)
	for i := 0; i < length; i++ {

		pre := i - 1
		next := i + 1

		palStr = fmt.Sprintf("%c", s[i])

		for pre >= 0 {
			if s[pre] == s[i] {
				palStr = fmt.Sprintf("%c", s[pre]) + palStr
				pre--

				continue
			}

			break
		}

		for next < length {
			if s[next] == s[i] {
				palStr = fmt.Sprintf("%c", s[next]) + palStr
				next++

				continue
			}

			break
		}

		for pre >= 0 && next < length {
			if s[pre] != s[next] {
				break
			}

			palStr = fmt.Sprintf("%c", s[pre]) + palStr + fmt.Sprintf("%c", s[next])
			pre--
			next++
		}

		if len(result) <= len(palStr) {
			result = palStr
		}
	}

	return result
}

func main() {
	fmt.Println(longestPalindrome("baban"))
}
