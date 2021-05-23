/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
*/

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ret := 0
	treeQueue := make([]*TreeNode, 0)
	treeQueue = append(treeQueue, root)
	for len(treeQueue) > 0 {
		ret++
		curNodes := len(treeQueue)
		for i := 0; i < curNodes; i++ {
			if treeQueue[i].Left != nil {
				treeQueue = append(treeQueue, treeQueue[i].Left)
			}
			if treeQueue[i].Right != nil {
				treeQueue = append(treeQueue, treeQueue[i].Right)
			}
		}
		treeQueue = treeQueue[curNodes:]
	}
	return ret
}

// @lc code=end

