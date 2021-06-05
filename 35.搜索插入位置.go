/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	return searchLess(nums, target)
}

func searchLess(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + ((right - left) >> 1)
		if nums[mid] < target {
			if (mid == len(nums)-1) || (nums[mid+1] >= target) {
				return mid + 1
			}
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return 0
}

// @lc code=end

