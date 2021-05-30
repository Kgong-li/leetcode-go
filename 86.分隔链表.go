/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return head
	}
	lowList, highList := &ListNode{Next: nil}, &ListNode{Next: nil}
	lowPre, highPre := lowList, highList

	for head != nil {
		if head.Val < x {
			lowPre.Next = head
			lowPre = lowPre.Next
		} else {
			highPre.Next = head
			highPre = highPre.Next
		}
		head = head.Next
	}
	highPre.Next = nil
	lowPre.Next = highList.Next
	return lowList.Next
}

// @lc code=end

