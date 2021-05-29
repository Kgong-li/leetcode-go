/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
/*
func containsNearbyDuplicate(nums []int, k int) bool {
	len := len(nums)
	if len < 2 {
		return false
	}

	for i := 0; i < len; i++ {
		left := i + 1
		for left <= i+k && left < len {
			if nums[i] == nums[left] {
				return true
			}
			left++
		}
	}
	return false
}
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 2 || k <= 0 {
		return false
	}

	record := make(map[int]bool, 0)
	for i, n := range nums {
		if _, ok := record[n]; ok {
			return true
		}

		record[n] = true
		if len(record) == k+1 {
			delete(record, nums[i-k])
		}
	}

	return false
}

// @lc code=end

