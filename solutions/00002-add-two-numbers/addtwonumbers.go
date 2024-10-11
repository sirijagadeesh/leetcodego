package addtwonumbers

// ListNode for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	num1, num2, carry, current := 0, 0, 0, head

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 == nil {
			num1 = 0
		} else {
			num1 = l1.Val
			l1 = l1.Next
		}

		if l2 == nil {
			num2 = 0
		} else {
			num2 = l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: (num1 + num2 + carry) % 10}
		carry = (num1 + num2 + carry) / 10
		current = current.Next

	}

	return head.Next
}
