package main

import "fmt"

// 罗马数字转int
func romanToInt(s string) int {

	target := 0
	length := len(s)

	RomanMap := map[uint8]int{
		73: 1,
		86: 5,
		88: 10,
		76: 50,
		67: 100,
		68: 500,
		77: 1000,
	}

	for i := 0; i < length-1; i++ {
		if RomanMap[s[i]] < RomanMap[s[i+1]] {
			target -= RomanMap[s[i]]
		} else {
			target += RomanMap[s[i]]
		}
	}

	target += RomanMap[s[length-1]]

	return target
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
