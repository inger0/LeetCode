package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{0, nil}
	tmpResult := &result

	for true {
		if l1 == nil && l2 == nil {
			break
		}

		if l1 == nil && l2 != nil {
			tmpResult.Next = l2
			break
		}

		if l2 == nil && l1 != nil {
			tmpResult.Next = l1
			break
		}

		if l1.Val <= l2.Val {
			tmpResult.Next = l1
			l1 = l1.Next
			tmpResult = tmpResult.Next
		} else {
			tmpResult.Next = l2
			l2 = l2.Next
			tmpResult = tmpResult.Next
		}
	}

	return result.Next
}

func main() {

}
