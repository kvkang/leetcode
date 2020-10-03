package main

/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
func romanToInt(s string) int {
	return romanToInt2(s)
}

func romanToInt2(s string) int {
	var sum int
	chValue := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); {
		iv := chValue[s[i]]

		var j = i + 1
		for j < len(s) {
			if s[i] != s[j] {
				break
			}
			j++
		}

		if j == len(s) {
			sum = sum + (j-i)*iv
			break
		}

		jv := chValue[s[j]]

		if jv > iv {
			sum = sum + jv - iv
			i = j + 1
		} else if jv <= iv {
			sum = sum + (j-i)*iv
			i = j
		}
	}
	return sum
}

//func romanToIntSelfSolution(s string) int {
//	var sum int
//	for i := 0; i < len(s); {
//		if s[i] == 'I' {
//			for j := i + 1; j < len(s); j++ {
//				if s[j] == 'V' {
//					i++
//					sum = sum + 3
//					break
//				} else if s[j] == 'X' {
//					i++
//					sum = sum + 8
//					break
//				} else if s[j] == 'I' {
//					i++
//					sum = sum + 1
//					continue
//				} else {
//					break
//				}
//			}
//			i++
//			sum = sum + 1
//		} else if s[i] == 'V' {
//			i++
//			sum = sum + 5
//		} else if s[i] == 'X' {
//			for j := i + 1; j < len(s); j++ {
//				if s[j] == 'L' {
//					i++
//					sum = sum + 30
//					break
//				} else if s[j] == 'C' {
//					i++
//					sum = sum + 80
//					break
//				} else if s[j] == 'X' {
//					i++
//					sum = sum + 10
//					continue
//				} else {
//					break
//				}
//			}
//			i++
//			sum = sum + 10
//		} else if s[i] == 'L' {
//			i++
//			sum = sum + 50
//		} else if s[i] == 'C' {
//			for j := i + 1; j < len(s); j++ {
//				if s[j] == 'D' {
//					i++
//					sum = sum + 300
//					break
//				} else if s[j] == 'M' {
//					i++
//					sum = sum + 800
//					break
//				} else if s[j] == 'C' {
//					i++
//					sum = sum + 100
//					continue
//				} else {
//					break
//				}
//			}
//			i++
//			sum = sum + 100
//		} else if s[i] == 'D' {
//			i++
//			sum = sum + 500
//		} else if s[i] == 'M' {
//			i++
//			sum = sum + 1000
//		} else {
//			i++
//		}
//	}
//	return sum
//}

// @lc code=end
