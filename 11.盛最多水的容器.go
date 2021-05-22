/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	max, left, right := 0, 0, len(height)-1
	for left < right {
		width := right - left
		high := 0
		if height[left] > height[right] {
			high = height[right]
			right--
		} else {
			high = height[left]
			left++
		}

		temp := high * width
		if temp > max {
			max = temp
		}
	}
	return max
}

// @lc code=end

