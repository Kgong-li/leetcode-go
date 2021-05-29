/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	l := len(s)
	if l == 0 {
		return
	}
	left, right := 0, l-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// @lc code=end

