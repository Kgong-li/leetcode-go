/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	ns := make(map[int]int, 0)
	for i, num := range nums {
		if _, ok := ns[target-num]; ok {
			return []int{i, ns[target-num]}
		}
		ns[num] = i
	}
	return nil
}

// @lc code=end

