/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	nePos := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			if i != nePos {
				nums[i], nums[nePos] = nums[nePos], nums[i]
			}
			nePos++
		}
	}
	return nePos
}

// @lc code=end

