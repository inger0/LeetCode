package main

//Given a linked list, swap every two adjacent nodes and return its head.
//
//You may not modify the values in the list's nodes, only nodes itself may be changed.
//
//
//
//Example:
//
//Given 1->2->3->4, you should return the list as 2->1->4->3.

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
