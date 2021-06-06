/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return [][]int{}
	}

	c, res := []int{}, [][]int{}
	sort.Ints(candidates)
	findCombinationSum(candidates, target, 0, c, &res)
	return res
}

func findCombinationSum(nums []int, target, index int, c []int, res *[][]int) {
	if target <= 0 {
		if target == 0 {
			tmp := make([]int, len(c))
			copy(tmp, c)
			*res = append(*res, tmp)
		}
		return
	}

	for i := index; i < len(nums); i++ {
		if nums[i] > target {
			break
		}
		c = append(c, nums[i])
		findCombinationSum(nums, target-nums[i], i, c, res)
		c = c[:len(c)-1]
	}
}

// @lc code=end

