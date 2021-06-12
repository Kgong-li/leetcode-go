/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	matrix, down, up := make([][]byte, numRows, numRows), 0, numRows-2
	for i := 0; i < len(s); {
		if down < numRows {
			matrix[down] = append(matrix[down], byte(s[i]))
			down++
			i++
		} else if up > 0 {
			matrix[up] = append(matrix[up], byte(s[i]))
			up--
			i++
		} else {
			down = 0
			up = numRows - 2
		}
	}

	res := make([]byte, 0, len(s))
	for _, row := range matrix {
		for _, v := range row {
			res = append(res, v)
		}
	}

	return string(res)
}

// @lc code=end

