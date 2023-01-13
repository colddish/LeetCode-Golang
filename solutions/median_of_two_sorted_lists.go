package solutions

/*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

The overall run time complexity should be O(log (m+n)).
*/

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	target := (len(nums1) + len(nums2)) / 2
	if (len(nums1)+len(nums2))%2 != 0 {
		return float64(findMedian(target, nums1, nums2))
	} else {
		median1 := findMedian(target, nums1, nums2)
		median2 := findMedian(target-1, nums1, nums2)
		return float64(median1+median2) / 2.0
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func findMedian(target int, nums1, nums2 []int) int {
	//Base Cases
	if len(nums1) == 0 {
		return nums2[target]
	} else if len(nums2) == 0 {
		return nums1[target]
	} else if target == 0 {
		return min(nums1[0], nums2[0])
	} else if target == len(nums1)+len(nums2)-1 {
		return max(nums1[len(nums1)-1], nums2[len(nums2)-1])
	} else if target == 1 {
		if nums1[0] > nums2[0] {
			if len(nums2) < 2 {
				return nums1[0]
			}
			return min(nums1[0], nums2[1])
		} else {
			if len(nums1) < 2 {
				return nums2[0]
			}
			return min(nums1[1], nums2[0])
		}
	}
	mid1, mid2 := min(len(nums1)/2, target), min(len(nums2)/2, target)
	if target <= (len(nums1)+len(nums2))/2 {
		if nums1[mid1] < nums2[mid2] {
			if target == len(nums1)+len(nums2[:mid2]) {
				return nums2[mid2]
			} else if len(nums1) > mid1+1 && nums2[mid2] < nums1[mid1+1] && len(nums1[:mid1+1])+len(nums2[:mid2+1]) == target+1 {
				return nums2[mid2]
			}
			return findMedian(target, nums1, nums2[:mid2])
		} else if nums1[mid1] > nums2[mid2] {
			if target == len(nums1[:mid1])+len(nums2) {
				return nums1[mid1]
			} else if len(nums2) > mid2+1 && nums1[mid1] < nums2[mid2+1] && len(nums1[:mid1+1])+len(nums2[:mid2+1]) == target+1 {
				return nums1[mid1]
			}
			return findMedian(target, nums1[:mid1], nums2)
		} else {
			if target+1 == len(nums1[:mid1+1])+len(nums2[:mid2+1]) || target+2 == len(nums1[:mid1+1])+len(nums2[:mid2+1]) {
				return nums1[mid1]
			}
			return findMedian(target, nums1[:mid1+1], nums2[:mid2+1])
		}
	} else {
		if nums1[mid1] < nums2[mid2] {
			if mid1 == 0 {
				if mid2 == target && nums1[mid1] > nums2[mid2] {
					return nums1[mid1]
				}
				return findMedian(target-1, nums1[:mid1], nums2)
			}
			return findMedian(target-mid1, nums1[mid1:], nums2)
		} else if nums1[mid1] > nums2[mid2] {
			if mid2 == 0 {
				if mid1 == target && nums2[mid2] > nums1[mid1] {
					return nums2[mid2]
				}
				return findMedian(target-1, nums1, nums2[:mid2])
			}
			return findMedian(target-mid2, nums1, nums2[mid2:])
		} else {
			if len(nums1) > len(nums2) {
				return findMedian(target-mid1, nums1[mid1:], nums2)
			} else {
				if mid2 == 0 {
					return findMedian(target-1, nums1, nums2[:mid2])
				}
				return findMedian(target-mid2, nums1, nums2[mid2:])
			}
		}
	}
}
