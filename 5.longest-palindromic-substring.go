package main

/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 */

// @lc code=start
func longestPalindrome(s string) string {
	// return longestPalindromeSelfSolution(s)
	return longestPalindromeApproach1(s)
	// return longestPalindromeApproach2(s)
	// return longestPalindromeApproach3(s)
	// return longestPalindromeApproach4(s) // same as selfSolution
	// return longestPalindromeApproach5(s)
}

func isLongestCommonSubstring(s string, x, y int) bool {
	for ; x < y; x, y = x+1, y-1 {
		if s[x] != s[y] {
			return false
		}
	}
	return true
}

/*
 * Self solution
 */
func longestPalindromeSelfSolution(s string) string {
	/*
	 * - longest common substring
	 * - for s Palindrome [i, j]
	 *     s[i] == s[j], s[i + 1] == s[j - 1] ... s[i + ?] == s[j - ?]
	 *     and (i+j) / 2 == 0 or (i+j) %2 == 1
	 *     so, set for an palindrome string s[i:j], set mid = i + j, mid >=i,
	 *        got s[mid/1] == s[(mid+0.5)/1] , s[mid/1 - 1] == s[(mid+0.5)/1 + 1] ... s[i] == s[j]
	 *   use all of the index of s as mid to find palindrome, find the max len of palindrome
	 */

	var (
		bs  = []byte(s)
		sub string
	)

	for i := float64(len(bs) - 1); i >= 0; i -= 0.5 {
		var x, y, step int
		for a, b := int(i/1), int(i+0.5/1); a >= 0 && b <= len(bs)-1 && bs[a] == bs[b]; a, b = a-1, b+1 {
			x, y, step = a, b, step+1
		}
		if step > 0 {
			if l := y - x + 1; l > len(sub) {
				sub = s[x : y+1]
			}
		}
	}

	return sub
}

/*
 * Approach 1:
 * - Reverse S and become S', Find the longest common substring between S and S', which must also be the longest palindromic substring.
 */
func longestPalindromeApproach1(s string) string {
	if len(s) <= 1 {
		return s
	}

	var l, r, b, e int
	for r < len(s) {
		for r+1 < len(s) && s[r+1] == s[l] {
			r++
		}

		for l-1 >= 0 && r+1 < len(s) && s[l-1] == s[r+1] {
			l--
			r++
		}

		if (r - l) > (e - b) {
			e = r
			b = l
		}

		l = (l+r)/2 + 1
		r = l
	}
	return s[b : e+1]
}

/*
 * Approach 2: Brute Force
 * - check all of sub string
 * Time complexity : O(n^3)
 * Space complexity : O(1)
 */
func longestPalindromeApproach2(s string) string {
	var maxLen, start, end int
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			// s[i,j] ---> all of sub string
			if isLongestCommonSubstring(s, i, j) {
				if l := j - i + 1; l > maxLen {
					maxLen, start, end = l, i, j+1
				}
			}
		}
	}
	return s[start:end]
}

/*
 * Approach 3: Dynamic Programming(动态规划)
 * - an palindrome string s[i:j], it's sub string s[i+1:j-1] is palindrome
 * Time complexity : O(n^2)
 * Space complexity : O(n^2)
 */
func longestPalindromeApproach3(s string) string {
	// 离散数学思维
	/*
	 * s[i: j] has two result, one is palindrome, other is not.
	 * for palindrome s[i:j] --> s[i + 1: j - 1] && s[i] == s[j]
	 * and palindrome s[i:j] --> s[(i+j)/2] == s[(i+j+1)/2]
	 * so mark bound site, l = j - i
	 *   when l == 0, s[i:i] is palindrome
	 *   when l == 1 and s[i] == s[j], s[i, j] is palindrome
	 */
	n := len(s)
	ans := ""
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for l := 0; l < n; l++ { // l == j - i
		for i := 0; i+l < n; i++ {
			j := i + l
			if l == 0 {
				// s[i:i] is palindrome
				dp[i][j] = 1
			} else if l == 1 {
				if s[i] == s[j] {
					// s[i:i+1] is palindrome when s[i] == s[j]
					dp[i][j] = 1
				}
			} else {
				if s[i] == s[j] {
					// if s[i+1, j-1] is palindrome and s[i] == s[j], s[i:j] is palindrome
					dp[i][j] = dp[i+1][j-1]
				}
			}
			if dp[i][j] > 0 && l+1 > len(ans) {
				ans = s[i : i+l+1]
			}
		}
	}
	return ans
}

/*
 * Approach 4: Expand Around Center(中心扩展算法)
 * - check all of sub string
 * Time complexity : O(n^2)
 * Space complexity : O(1)
 * is same as selfSolution, can see selfSolution comment
 */
func longestPalindromeApproach4(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

/*
 * Approach 5: Manacher
 * Time complexity : O(n^2)
 * Space complexity : O(1)
 * use expand around center and Dynamic Programming
 * Accepted
 * 103/103 cases passed (0 ms)
 * Your runtime beats 100 % of golang submissions
 * Your memory usage beats 38.46 % of golang submissions (3.2 MB)
 */
func longestPalindromeApproach5(s string) string {
	/*
	 * for a palindrome string dp[left: right], len(dp) % 2 == 1
	 *   Note:
	 *       c as dp[left : right]' s center
	 *       p[c] as the index == c, the center longest expand len
	 *   So: 2c == (right + left) , p[c] = (right - left) / 2 = right - c
	 * for the palindrome string which center index is i,
	 *   when c < i <= right,
	 *     Note:
	 *        i' is the mirror of i when the center c -->  c - i' = i - c --> i' = 2c - i
	 *        p[i] as center index == i, the center longest expand len
	 *        p[i'] as center index == i', the center longest expand len
	 *     now think of relationship between p[i] with p[i']
	 *        if left < (i' - p[i']),
	 *           So：
	 *              s[i' - p[i'] - 1] != s[i' + p[i'] + 1] --> s[i + p[i'] + 1] != s[i - p[i'] - 1]
	 *              dp[i - p[i'] : i + p[i']] is the mirror string of dp[i' - p[i'] when center index == c : i + p[i']] --> dp[i - p[i'] : i + p[i']] is palindrome string
	 *           So:
	 *              p[i] == p[i']
	 *              And left < (i' - p[i']) , i' = 2c - i
	 *           --->
	 *              if i + p[i'] < right ---> p[i] == p[i']
	 *        if left == (i' - p[i'])
	 *           as left < (i' - p[i']) we can get dp[i - p[i'] : i + p[i']] is palindrome string
	 *           and used expand Around Center to checkout dp[i - p[i'] - x : i + p[i'] + x]
	 *           So p[i] >= p[i']
	 *              And left < (i' - p[i']) , i' = 2c - i
	 *           --->
	 *              if i + p[i'] == right ---> p[i] >= p[i'], p[i] real value need use expand Around center to checkout
	 *        if left > (i' - p[i'])
	 *           Note: l , l = i' - left. l < p[i']
	 *           So:
	 *              l = i' - left = right - i
	 *              dp[i'-l: i'+l] is a palindrome string, it's mirros string dp[i-l: i+l] is also palindrome string
	 *              s[right + 1] == s[i+l+1]
	 *              s[left - 1] == s[i' + l + 1] == s[i- l - 1]
	 *           And: s[left - 1] != s[right + 1] ---> s[i + l + 1] != s[i-l-1]
	 *           So: p[i] = l
	 *           ---->
	 *              if i + p[i'] > right p[i] = right - i
	 *           so p[i] = right - i
	 *   when i > right
	 *       get p[i] used expand Around Center
	 */

	var (
		newS              string                    // an string which it's sub palindrome string's len % 2 = 1, and len(newS) == 2 * len(s) + 1
		expandLen         = make([]int, 2*len(s)+1) // index is newS's index and the palindrome string's center index, stone the palindrome string righet long (expand long of center)
		maxExpandLenIndex = 0                       // stone the currunt longest palindrome string we find from newS
		curCenter         = 0                       // for the curMaxRight's center
		curMaxRight       = 0
	)
	{
		// build newS
		// to make s's palindrome string dp[x:y] to new s's palindrome string dp[2x+1: 2y+1]
		// so: dp[2x+1: 2y+1] can len(dp[2x+1:2y+1]) % 2 == 1 --> ((2y+1) - (2x+1) + 1) %2 == (2y - 2x + 1) % 2 == 1
		helpBytes := make([]byte, 0, 2*len(s)+1)
		helpBytes = append(helpBytes, '#')
		for i := 0; i < len(s); i++ {
			helpBytes = append(helpBytes, s[i], '#')
		}
		newS = string(helpBytes)
	}

	for i := 0; i < len(newS); i++ {
		if i <= curMaxRight {
			iMirror := 2*curCenter - i
			if expandLen[iMirror]+i < curMaxRight {
				expandLen[i] = expandLen[iMirror]
			} else if expandLen[iMirror]+i > curMaxRight {
				expandLen[i] = curMaxRight - i
			} else {
				// expandLen[iMirror]+i == curRight
				expandLen[i] = expand(newS, i-expandLen[iMirror]-1, i+expandLen[iMirror]+1)
			}
		} else {
			expandLen[i] = expand(newS, i-1, i+1)
		}

		if expandLen[i] > expandLen[maxExpandLenIndex] {
			maxExpandLenIndex = i
		}

		if nowRight := i + expandLen[i]; nowRight > curMaxRight {
			curCenter = i
			curMaxRight = nowRight
		}
	}

	start := (maxExpandLenIndex - expandLen[maxExpandLenIndex]) / 2
	end := (maxExpandLenIndex + expandLen[maxExpandLenIndex]) / 2
	return s[start:end]
}

func expand(s string, left, right int) int {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return (right - left - 2) / 2
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end
