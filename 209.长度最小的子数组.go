/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	len := len(nums)
	if len == 0 {
		return 0
	}
	left, right, sum, ret := 0, -1, 0, len+1
	for left < len {
		if right+1 < len && sum < target {
			right++
			sum += nums[right]
		} else {
			sum -= nums[left]
			left++
		}

		if sum >= target {
			ret = min(ret, right-left+1)
		}
	}

	if ret == len+1 {
		return 0
	}
	return ret
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

