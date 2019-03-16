package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	//ListNodeMap := make(map[int]int)
	//index := 0
	//for head.Next != nil {
	//	ListNodeMap[index] = head.Val
	//	head = head.Next
	//	index++
	//}
	//
	//ListNodeMap[index] = head.Val
	//var result *ListNode
	//
	//target := index - n + 1
	//for index >= 0 {
	//	switch true {
	//	case index == target:
	//		break
	//
	//	default:
	//		if result == nil {
	//			result = &ListNode{ListNodeMap[index], nil}
	//		} else {
	//			tmp := &ListNode{ListNodeMap[index], result}
	//			result = tmp
	//		}
	//	}
	//
	//	index--
	//}
	//
	//return result

	result := ListNode{0, nil}
	result.Next = head

	length := 0
	first := head
	for true {
		length++
		if first.Next == nil {
			break
		}
		first = first.Next
	}

	length -= n

	first = &result
	for length > 0 {
		first = first.Next
		length--
	}

	first.Next = first.Next.Next

	return result.Next
}

func main() {

}
