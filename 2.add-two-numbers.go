package main

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		carry int
		head  ListNode
		cur   = &head
	)

	for l1 != nil || l2 != nil {
		var l1Val, l2Val int
		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}

		sum := carry + l1Val + l2Val
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10, Next: nil}
		cur = cur.Next
	}

	if carry != 0 {
		cur.Next = &ListNode{Val: carry, Next: nil}
	}

	return head.Next
}

// @lc code=end
