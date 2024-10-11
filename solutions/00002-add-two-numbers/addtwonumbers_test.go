package addtwonumbers

import (
	"reflect"
	"testing"
)

func NewSingleLinkedList(arr []int) *ListNode {
	head := &ListNode{}
	cur := head

	for _, j := range arr {
		cur.Next = &ListNode{Val: j}
		cur = cur.Next
	}
	return head.Next
}

func Test_addTwoNumbers(t *testing.T) {
	testDatas := []struct {
		name     string
		arg1     *ListNode
		arg2     *ListNode
		expected *ListNode
	}{
		{
			name:     "one",
			arg1:     NewSingleLinkedList([]int{2, 4, 3}),
			arg2:     NewSingleLinkedList([]int{5, 6, 4}),
			expected: NewSingleLinkedList([]int{7, 0, 8}),
		},
		{
			name:     "two",
			arg1:     NewSingleLinkedList([]int{5}),
			arg2:     NewSingleLinkedList([]int{5}),
			expected: NewSingleLinkedList([]int{0, 1}),
		},
		{
			name:     "three",
			arg1:     NewSingleLinkedList([]int{1}),
			arg2:     NewSingleLinkedList([]int{9, 9, 9}),
			expected: NewSingleLinkedList([]int{0, 0, 0, 1}),
		},
		{
			name:     "four",
			arg1:     NewSingleLinkedList([]int{9, 9, 9}),
			arg2:     NewSingleLinkedList([]int{1}),
			expected: NewSingleLinkedList([]int{0, 0, 0, 1}),
		},
		{
			name:     "five",
			arg1:     NewSingleLinkedList([]int{4, 3, 1}),
			arg2:     NewSingleLinkedList([]int{1}),
			expected: NewSingleLinkedList([]int{5, 3, 1}),
		},
		{
			name:     "six",
			arg1:     NewSingleLinkedList([]int{1}),
			arg2:     NewSingleLinkedList([]int{4, 3, 1}),
			expected: NewSingleLinkedList([]int{5, 3, 1}),
		},
	}

	for _, testData := range testDatas {
		t.Run(testData.name, func(t *testing.T) {
			if result := addTwoNumbers(testData.arg1, testData.arg2); !reflect.DeepEqual(result, testData.expected) {
				t.Errorf("expected %v, got %v", testData.expected, result)
			}
		})
	}
}
