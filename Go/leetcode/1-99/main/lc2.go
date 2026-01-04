package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// runtime: O(N + M)
// space: O(max(N + M))
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := new(ListNode)
	curr := dummy
	sum := 0
	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{sum % 10, nil}
		curr = curr.Next
		sum /= 10
	}
	if sum != 0 {
		curr.Next = &ListNode{sum, nil}
	}
	return dummy.Next
}
