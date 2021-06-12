/*
 * @lc app=leetcode.cn id=8 lang=golang
 *
 * [8] 字符串转换整数 (atoi)
 */

// @lc code=start
func myAtoi(s string) int {
	head, signed, space, sign, digits := false, true, true, 1, []int{}
	for _, c := range s {
		if c == ' ' && space {
			continue
		}
		if signed {
			if c == '+' {
				signed = false
				space = false
				continue
			} else if c == '-' {
				sign = -1
				signed = false
				space = false
				continue
			}
		}

		if c < '0' || c > '9' {
			break
		}

		if !head && c == '0' {
			space, signed = false, false
			continue
		}
		space, signed, head = false, false, true
		digits = append(digits, int(c-48))
	}

	var num, max, place int64 = 0, 0, 1
	if sign > 0 {
		max = math.MaxInt32
	} else {
		max = 0 - math.MinInt32
	}
	for i := len(digits) - 1; i >= 0; i-- {
		num += int64(digits[i]) * place
		place *= 10
		if num > max || (place > max && i > 0) {
			return int(max * int64(sign))
		}
	}

	return int(num * int64(sign))
}

// @lc code=end
