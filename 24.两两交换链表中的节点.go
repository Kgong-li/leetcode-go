/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nHead := &ListNode{Val: 0, Next: head}
	pre := nHead
	lNode := pre.Next
	rNode := lNode.Next
	for lNode != nil && rNode != nil {
		pre.Next = rNode
		lNode.Next = rNode.Next
		rNode.Next = lNode

		pre = lNode
		lNode = lNode.Next
		if lNode != nil {
			rNode = lNode.Next
		}
	}
	return nHead.Next
}

// @lc code=end

