/*
 * @lc app=leetcode.cn id=41 lang=golang
 *
 * [41] 缺失的第一个正数
 */

// @lc code=start
func firstMissingPositive(nums []int) int {
	nMap := make(map[int]int, len(nums))
	for _, num := range nums {
		nMap[num] = num
	}

	for index := 1; index < len(nums)+1; index++ {
		if _, ok := nMap[index]; !ok {
			return index
		}
	}
	return len(nums) + 1
}

// @lc code=end

