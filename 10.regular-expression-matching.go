package main

/*
 * @lc app=leetcode id=10 lang=golang
 *
 * [10] Regular Expression Matching
 */

// @lc code=start
func isMatch(s string, p string) bool {
	// return isMatchSelfSolution([]byte(s), []byte(p))
	//	return isMatch([]byte(s), []byte(p))
	return false
}

//type chItem struct {
//	num      int
//	numIsMin bool
//	ch       byte
//}

//
//func buildItems(s []byte) []*chItem {
//	var lastItem *chItem
//	var sItems = make([]*chItem, 0, len(s))
//
//	for i := 0; i < len(s); i++ {
//		if lastItem == nil {
//			lastItem = &chItem{ch: s[i], num: 1}
//			sItems = append(sItems, lastItem)
//		} else if s[i] == '*' {
//			if !lastItem.numIsMin {
//				lastItem.numIsMin = true
//				lastItem.num = lastItem.num - 1
//			}
//		} else if lastItem.ch != s[i] {
//			lastItem = &chItem{ch: s[i], num: 1}
//			sItems = append(sItems, lastItem)
//		} else {
//			lastItem.num = lastItem.num + 1
//		}
//	}
//	return sItems
//}
//
//func newItem(ch byte, num int, isMin bool, rest []*chItem) []*chItem {
//	if num == 0 && !isMin {
//		return rest
//	}
//
//	items := make([]*chItem, 0, len(rest)+1)
//	items = append(items, &chItem{num, isMin, ch})
//	items = append(items, rest...)
//	return items
//}
//
//func itemMatch(s, p []*chItem) bool {
//	if len(p) == 0 {
//		return len(s) == 0
//	} else if len(s) == 0 {
//		if p[0].num == 0 {
//			return itemMatch(s, p[1:])
//		}
//		return false
//	}
//
//	if s[0].num < p[0].num {
//		if p[0].ch == '.' {
//			return itemMatch(
//				s[1:],
//				newItem(p[0].ch, p[0].num-s[0].num, p[0].numIsMin, p[1:]),
//			)
//		}
//		return false
//	} else if s[0].num == p[0].num {
//		if p[0].ch == '.' || p[0].ch {
//			return itemMatch(
//				s[1:],
//				newItem(p[0].ch, 0 /* p[0].num-s[0].num */, p[0].numIsMin, p[1:]),
//			)
//		}
//		if p[0].ch == s[0].ch {
//			return itemMatch(s[1:], p[1:])
//		}
//		return false
//	} else {
//		if p[0].ch == '.' {
//			if p[0].numIsMin {
//				if itemMatch(s[1:], newItem(p[0].ch, 0, p[0].numIsMin, p[1:])) {
//					return true
//				}
//			}
//			return itemMatch(
//				newItem(s[0].ch, s[0].num-p[0].num, s[0].numIsMin, s[1:]),
//				p[1:],
//			)
//		}
//
//		if p[0].ch == s[0].ch {
//			if p[0].numIsMin {
//				if itemMatch(s[1:], newItem(p[0].ch, 0, p[0].numIsMin, p[1:])) {
//					return true
//				}
//			}
//
//			return itemMatch(
//				newItem(s[0].ch, s[0].num-p[0].num, s[0].numIsMin, s[1:]),
//				p[1:],
//			)
//		}
//
//		if p[0].num == 0 {
//			return itemMatch(s, p[1:])
//		}
//	}
//	return false
//}
//
//func isMatchSelfSolution(s, p []byte) bool {
//	return itemMatch(buildItems(s), buildItems(p))
//}

// @lc code=end
