/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] 两个数组的交集
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		return intersection(nums2, nums1)
	}

	ret := make([]int, 0)
	n1map := make(map[int]struct{})

	for _, num := range nums1 {
		n1map[num] = struct{}{}
	}

	for _, num := range nums2 {
		if _, ok := n1map[num]; ok {
			ret = append(ret, num)
			delete(n1map, num)
		}
	}
	return ret
}

// @lc code=end

