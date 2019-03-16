package main

import "fmt"

//Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
//For example, given n = 3, a solution set is:
//
//[
//  "((()))",
//  "(()())",
//  "(())()",
//  "()(())",
//  "()()()"
//]

func generateParenthesis(n int) []string {
	var ans []string
	return backtrack(ans, "", 0, 0, n)
}

func backtrack(ans []string, cur string, open int, close int, max int) []string {

	if len(cur) == max*2 {
		ans = append(ans, cur)
		return ans
	}

	if open < max {
		ans = backtrack(ans, cur+"(", open+1, close, max)
	}

	if close < open {
		ans = backtrack(ans, cur+")", open, close+1, max)
	}

	return ans
}

func main() {
	fmt.Println(generateParenthesis(4))
}
