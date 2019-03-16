package main

import "fmt"

//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//An input string is valid if:
//
//Open brackets must be closed by the same type of brackets.
//Open brackets must be closed in the correct order.
//Note that an empty string is also considered valid.

// 栈的思想
func isValid(s string) bool {
	var stack []string
	digitMap := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, v := range s {
		strV := fmt.Sprintf("%c", v)
		if len(stack) == 0 {
			stack = append(stack, strV)
			continue
		}

		if digitMap[strV] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, strV)
		}
	}

	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("({)}[]"))
}
