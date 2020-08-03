package main

/*
 * @lc app=leetcode id=1 lang=golang
 *
 * [1] Two Sum
 */

// @lc code=start

func twoSum(nums []int, target int) []int {
	if len(nums) >= 2 {
		has := make(map[int]int, len(nums))
		for ai, a := range nums {
			b := target - a
			if bi, ok := has[b]; ok {
				return []int{bi, ai}
			}
			has[a] = ai
		}
	}
	return nil
}

// @lc code=end
