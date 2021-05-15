package main

import "fmt"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)

	k := (len1 + len2 + 1) / 2
	isOdd := false
	if (len1+len2)%2 == 1 {
		isOdd = true
	}
	median := findKthSortedArrays(nums1, nums2, k)
	if isOdd {
		return float64(median)
	}
	median1 := findKthSortedArrays(nums1, nums2, k+1)
	return float64(median+median1) / 2
}

func findKthSortedArrays(nums1, nums2 []int, k int) int {
	len1 := len(nums1)
	len2 := len(nums2)
	idx1 := 0
	idx2 := 0
	for {
		if idx1 == len1 {
			return nums2[idx2+k-1]
		}
		if idx2 == len2 {
			return nums1[idx1+k-1]
		}
		if k == 1 {
			return min(nums1[idx1], nums2[idx2])
		}
		half := k / 2
		m1 := min(len1, idx1+half)
		m2 := min(len2, idx2+half)
		if nums1[m1-1] < nums2[m2-1] {
			step := m1 - idx1
			idx1 += step
			k -= step
		} else {
			step := m2 - idx2
			idx2 += step
			k -= step
		}
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

func main() {
	nums1 := []int{1, 3, 4}
	nums2 := []int{2, 5, 6}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
