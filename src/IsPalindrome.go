package main

import "fmt"

func isPalindrome(x int) bool {

	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}

	tmp := x
	target := 0
	for x > 0 {
		target = target*10 + (x % 10)

		x /= 10
	}

	if tmp == target {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isPalindrome(22121))
}
