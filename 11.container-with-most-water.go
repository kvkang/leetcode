package main

/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	var max int
	//	for i := 0; i < len(height)-1; i++ {
	//		if height[i] == 0 {
	//			continue
	//		}
	//		var minJ = max/height[i] + 1 + i
	//		for j := len(height) - 1; j >= minJ; j-- {
	//			if height[j] == 0 {
	//				continue
	//			}
	//			v := (j - i) * maxAreaGetMin(height[i], height[j])
	//			if v > max {
	//				max = v
	//				minJ = max / height[i]
	//			}
	//		}
	//	}
	return max
}

func maxAreaGetMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end
