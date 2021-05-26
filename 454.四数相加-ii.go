/*
 * @lc app=leetcode.cn id=454 lang=golang
 *
 * [454] 四数相加 II
 */

// @lc code=start
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	n1n2map := make(map[int]int, len(nums1)*len(nums2))
	ret := 0
	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			n1n2map[n1+n2]++
		}
	}

	for _, n3 := range nums3 {
		for _, n4 := range nums4 {
			ret += n1n2map[0-n3-n4]
		}
	}
	return ret
}

// @lc code=end

