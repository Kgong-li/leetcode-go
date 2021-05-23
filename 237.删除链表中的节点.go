/*
 * @lc app=leetcode.cn id=237 lang=golang
 *
 * [237] 删除链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	if node == nil {
		return
	}
	cur := node
	for cur.Next.Next != nil {
		cur.Val = cur.Next.Val
		cur = cur.Next
	}
	cur.Val = cur.Next.Val
	cur.Next = nil
}

// @lc code=end

