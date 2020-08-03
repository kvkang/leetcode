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
	)

	for cur := &head; l1 != nil || l2 != nil || carry != 0; cur = cur.Next {

		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		cur.Next = &ListNode{Val: carry % 10, Next: nil}
		carry = carry / 10
	}
	return head.Next
}

// @lc code=end
