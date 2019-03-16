package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {

	result := ListNode{0, nil}
	result.Next = head
	tmpList := &result

	for true {
		if tmpList.Next == nil || tmpList.Next.Next == nil {
			break
		}

		firstNode := tmpList.Next
		secondNode := tmpList.Next.Next

		firstNode.Next = secondNode.Next
		secondNode.Next = firstNode
		tmpList.Next = secondNode

		tmpList = tmpList.Next.Next
	}

	return result.Next
}

func main() {

}
