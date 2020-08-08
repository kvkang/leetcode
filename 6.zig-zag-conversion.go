package main

/*
 * @lc app=leetcode id=6 lang=golang
 *
 * [6] ZigZag Conversion
 */

// @lc code=start
func convert(s string, numRows int) string {
	// return converSelfSolution(s, numRows)
	return converApproach1(s, numRows)
	//	return converApproach2(s, numRows)

}

/*
 * Self solution
 * 1158/1158 cases passed (4 ms)
 * Your runtime beats 92.24 % of golang submissions
 * Your memory usage beats 90.91 % of golang submissions (3.9 MB)
 */
func converSelfSolution(s string, numRows int) string {
	/*
	 * P(0 --> 0, 0)               A (5 --> 0, 2)  H   N
	 * A(1 --> 1, 0) P(4 --> 1, 1) L S I I G
	 * Y(2 --> 2, 0)               I   R
	 *
	 * P(0 --> 0, 0)                             I(6 --> 0, 3)    N
	 * A(1 --> 1, 0)               L(5 --> 1, 2) S  I G
	 * Y(2 --> 2, 0) A(4 --> 2, 1)               H R
	 * P(3 --> 3, 0)                             I
	 * ---->
	 * s[0] ....                                                    s[2*numRows - 3]
	 * s[1] ....                                  s[2*numRows - 4]
	 *  *
	 *  *                          s[numRows+1]
	 *  *              s[numRows]
	 * s[numRows - 1]
	 */

	if numRows <= 1 || len(s) <= numRows {
		return s
	}

	oneProcessLen := 2*numRows - 2
	bs := make([]byte, 0, len(s))
	for row, offset := 0, oneProcessLen; row < numRows; row, offset = row+1, offset-2 {
		hasOffset := offset != 0 && offset != oneProcessLen
		for j := row; j < len(s); j = j + oneProcessLen {
			bs = append(bs, s[j])
			if hasOffset {
				if newJ := j + offset; newJ < len(s) {
					bs = append(bs, s[newJ])
				} else {
					break
				}
			}
		}
	}
	return string(bs)
}

/*
 * Approach 1: Sort by Row
 * Time Complexity: O(n), where n == len(s)
 * Space Complexity: O(n)
 * 1158/1158 cases passed (0 ms)
 * Your runtime beats 100 % of golang submissions
 * Your memory usage beats 45.45 % of golang submissions (6.3 MB)
 */
func converApproach1(s string, numRows int) string {
	/*
	 * with the increment of index of s, append to the row index is move back and forth inside [0, numRows - 1]
	 */
	if numRows == 1 {
		return s
	}

	var rows [][]byte
	{
		if numRows > len(s) {
			rows = make([][]byte, len(s))
		} else {
			rows = make([][]byte, numRows)
		}
		colMaxValueLen := len(s)%numRows + 1
		for i := range rows {
			rows[i] = make([]byte, 0, colMaxValueLen)
		}
	}

	for i, curRow, goingDown := 0, 0, false; i < len(s); i++ {
		rows[curRow] = append(rows[curRow], s[i])

		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	bs := make([]byte, 0, len(s))
	for curRow := range rows {
		bs = append(bs, rows[curRow]...)
	}
	return string(bs)

}

/*
 * Approach 2: Sort by Row
 * Time Complexity: O(n), where n == len(s). Each index is visited once.
 * Space Complexity: O(n). For the cpp implementation, O(1). if return string is not considered extra space.
 * 1158/1158 cases passed (4 ms)
 * Your runtime beats 92.24 % of golang submissions
 * Your memory usage beats 81.82 % of golang submissions (3.9 MB)
 */
func converApproach2(s string, numRows int) string {
	/*
	 * like selfSolution
	 *   Visit all characters in row 0 first, then row 1, then row 2, and so on...
	 * For all whole numbers k,
	 *  Characters in row 0 are located at indexes k(2⋅numRows−2)
	 *  Characters in row numRows−1 are located at indexes k(2⋅numRows−2)+numRows−1
	 *  Characters in inner row ii are located at indexes k(2⋅numRows−2)+i and (k+1)(2⋅numRows−2)−i.
	 */
	if numRows == 1 {
		return s
	}
	bs := make([]byte, 0, len(s))
	cycleLen := 2*numRows - 2

	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j += cycleLen {
			bs = append(bs, s[j+i])
			if i != 0 && i != numRows-1 && j+cycleLen-i < len(s) {
				bs = append(bs, s[j+cycleLen-i])
			}
		}
	}
	return string(bs)
}

// @lc code=end
