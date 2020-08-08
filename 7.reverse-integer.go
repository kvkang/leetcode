package main

import "math"

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	// return reverseSelfSolution(x)
	return reverseApproach1(x)
}

func reverseSelfSolution(x int) (o int) {
	if x == 0 {
		return 0
	}

	lt0 := x < 0
	if lt0 {
		x = -x
	}

	for ; x%10 == 0; x = x / 10 {
	}

	for x > 0 {
		o = 10*o + (x % 10)
		x = x / 10
	}

	if lt0 {
		o = -o
	}

	if o < math.MinInt32 || o > math.MaxInt32 {
		o = 0
	}

	return
}

func reverseApproach1(x int) int {
	res := 0
	for x != 0 {
		carry := x % 10
		res = res*10 + carry
		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
		x /= 10
	}
	return res
}

// @lc code=end
