/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	res, left, right, maxleft, maxright := 0, 0, len(height)-1, 0, 0

	for left <= right {
		if height[left] <= height[right] {
			if height[left] > maxleft {
				maxleft = height[left]
			} else {
				res += maxleft - height[left]
			}
			left++
		} else {
			if height[right] >= maxright {
				maxright = height[right]
			} else {
				res += maxright - height[right]
			}
			right--
		}
	}

	return res
}

// @lc code=end

