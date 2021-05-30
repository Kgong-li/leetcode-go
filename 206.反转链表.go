/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var preNode, curNode *ListNode = nil, head
	for curNode != nil {
		nextNode := curNode.Next
		curNode.Next = preNode
		preNode = curNode
		curNode = nextNode
	}
	return preNode
}

// @lc code=end

