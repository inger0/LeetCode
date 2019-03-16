package main

import "fmt"

// 链表两数字相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var jinwei = 0
	var tmpArray []int

	for {
		if l1 == nil && l2 == nil {
			break
		}

		var sum = jinwei
		if l1 != nil {
			sum += l1.Val
		}

		if l2 != nil {
			sum += l2.Val
		}

		var nodeVal = sum % 10
		if sum >= 10 {
			jinwei = 1
		} else {
			jinwei = 0
		}

		tmpArray = append(tmpArray, nodeVal)

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}

	if jinwei > 0 {
		tmpArray = append(tmpArray, 1)
	}

	for i := len(tmpArray) - 1; i >= 0; i-- {
		if result == nil {
			result = &ListNode{tmpArray[i], nil}
		} else {

			tmp := &ListNode{tmpArray[i], result}
			result = tmp
		}
	}

	return result
}

func main() {

	var a1 = &ListNode{2, nil}
	var a2 = &ListNode{4, a1}
	var a3 = &ListNode{3, a2}
	a2.Next = a1
	a3.Next = a2

	var b1 = &ListNode{1, nil}
	var b2 = &ListNode{2, b1}
	var b3 = &ListNode{5, b2}
	b2.Next = b1
	b3.Next = b2

	fmt.Println(a3)

	var res = addTwoNumbers(a3, b3)
	fmt.Println(res)

	for true {
		if res == nil {
			break
		}

		fmt.Println(res.Val)

		res = res.Next
	}

}
