package main

/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	return 0
	// make sure that len(nums1) >= len(nums2)
	//	if len(nums1) > len(nums2) {
	//		return findMedianSortedArrays(nums2, nums1)
	//	}
	//
	//	var leftLen, rightLen int
	//	sumLen := len(nums1) + len(nums2)
	//	if sumLen%2 == 0 {
	//		leftLen = sumLen / 2
	//		rightLen = leftLen + 1
	//	} else {
	//		leftLen = sumLen / 2
	//		rightLen = leftLen
	//	}
	//
	//	var (
	//		start, end = math.MinInt64, math.MaxInt64
	//	)
	//
	//	for start < end || (len(nums1) == 0 && len(nums2) == 0) {
	//		if len(nums1) == 0 || len(nums2) == 0 {
	//			var nums []int
	//			if len(nums1) > 0 {
	//				nums = nums1
	//			} else {
	//				nums = nums2
	//			}
	//
	//			carry := len(nums) / 2
	//			start = nums[carry]
	//			if len(nums)%2 == 0 {
	//				end = nums[carry-1]
	//			} else {
	//				end = nums[carry]
	//			}
	//		}
	//	}
	//	return float64(start+end) / 2
}

// @lc code=end
