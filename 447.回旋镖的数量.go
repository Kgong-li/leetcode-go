/*
 * @lc app=leetcode.cn id=447 lang=golang
 *
 * [447] 回旋镖的数量
 */

// @lc code=start
func numberOfBoomerangs(points [][]int) int {
	ret := 0
	for i := 0; i < len(points); i++ {
		dict := make(map[int]int)
		for j := 0; j < len(points); j++ {
			d := dis(points[i], points[j])
			if v, ok := dict[d]; !ok {
				dict[d]++
			} else {
				ret += 2 * v
				dict[d]++
			}
		}
	}
	return ret
}

func dis(a, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}

// @lc code=end

