/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}

	r, w, b := 0, 0, 0
	for _, num := range nums {
		switch num {
		case 0:
			nums[b] = 2
			b++
			nums[w] = 1
			w++
			nums[r] = 0
			r++
		case 1:
			nums[b] = 2
			b++
			nums[w] = 1
			w++
		case 2:
			b++
		}
	}
	return
}

// @lc code=end

