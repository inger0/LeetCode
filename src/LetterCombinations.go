package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {

	NumberLetter := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	var target []string
	length := len(digits)

	return getCombination(0, digits, length, target, NumberLetter)
}

func getCombination(index int, digits string, length int, target []string, NumberLetter map[string]string) []string {

	if index == length {
		return target

	} else {
		digit := fmt.Sprintf("%c", digits[index])
		subStr := NumberLetter[digit]
		l := len(subStr)

		var tmpTarget []string

		for i := 0; i < l; i++ {
			tmp := fmt.Sprintf("%c", subStr[i])
			if len(target) == 0 {
				tmpTarget = append(tmpTarget, tmp)
			}

			for _, v := range target {
				tmpTarget = append(tmpTarget, v+tmp)
			}
		}

		target = tmpTarget

		return getCombination(index+1, digits, length, target, NumberLetter)
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
