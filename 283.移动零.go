/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	zero := 0
	for noZero := 0; noZero < len(nums); noZero++ {
		if nums[noZero] != 0 {
			if noZero != zero {
				nums[noZero], nums[zero] = nums[zero], nums[noZero]
				zero++
			} else {
				zero++
			}
		}
	}
	return
}

// @lc code=end

