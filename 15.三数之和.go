/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	result, left, right, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length-1; index++ {
		left, right = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			left = index - 1
		}

		for left < index && right > index {
			if left > 0 && nums[left] == nums[left-1] {
				left++
				continue
			}

			if right < length-1 && nums[right] == nums[right+1] {
				right--
				continue
			}

			addNum = nums[left] + nums[index] + nums[right]
			switch {
			case addNum == 0:
				result = append(result, []int{nums[left], nums[index], nums[right]})
				left++
				right--
			case addNum < 0:
				left++
			case addNum > 0:
				right--
			}
		}
	}
	return result
}

// @lc code=end

