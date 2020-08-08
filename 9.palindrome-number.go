package main

/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	return isPalindromeSelfSolution(x)
}

func isPalindromeSelfSolution(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	var reverseX int
	for i := x; reverseX < x; i = i / 10 {
		reverseX = 10*reverseX + i%10
	}

	return reverseX == x || reverseX/10 == x
}

// @lc code=end
