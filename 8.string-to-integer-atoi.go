package main

import "math"

/*
 * @lc app=leetcode id=8 lang=golang
 *
 * [8] String to Integer (atoi)
 */

// @lc code=start
func myAtoi(str string) int {
	return myAtoiSelfSolution(str)
}

func myAtoiSelfSolution(str string) (ret int) {
	i := 0
	var ch byte

	for ; i < len(str) && str[i] == ' '; i++ {
	}

	if i < len(str) && (str[i] == '-' || str[i] == '+') {
		ch = str[i]
		i++
	}

	for ; i < len(str) && str[i] >= '0' && str[i] <= '9'; i++ {
		ret = 10*ret + int(str[i]-'0')
		if ret > math.MaxInt32 {
			break
		}
	}

	if ch == '-' {
		ret = -ret
		if ret < math.MinInt32 {
			return math.MinInt32
		}
	} else if ret > math.MaxInt32 {
		return math.MaxInt32
	}

	return
}

// @lc code=end
