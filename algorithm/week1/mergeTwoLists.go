package main

type ListNode struct {
	Val int
	Next *ListNode
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	t := &ListNode{}
	head := t
	for l1 != nil || l2 != nil {
		if l1.Val > l2.Val {
			t.Next = l2
			l2 = l2.Next
		} else {
			t.Next = l1
			l1 = l1.Next
		}
		t = t.Next
		if l1 == nil {
			t.Next = l2
			break
		}
		if l2 == nil {
			t.Next = l1
			break
		}
	}
	return head.Next
}