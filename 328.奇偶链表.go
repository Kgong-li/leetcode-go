/*
 * @lc app=leetcode.cn id=328 lang=golang
 *
 * [328] 奇偶链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	oddList, evenList := &ListNode{Next: nil}, &ListNode{Next: nil}
	oddPre, evenPre := oddList, evenList
	index := 0
	for head != nil {
		if index%2 == 0 {
			oddPre.Next = head
			oddPre = head
		} else {
			evenPre.Next = head
			evenPre = head
		}
		head = head.Next
		index++
	}
	evenPre.Next = nil
	oddPre.Next = evenList.Next
	return oddList.Next
}

// @lc code=end

