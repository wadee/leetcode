package main

import "fmt"

const (
	UINT_MAX = ^uint(0)
	INT_MAX  = int(UINT_MAX >> 1)
	INT_MIN  = -INT_MAX - 1
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	len1 := len(nums1)
	len2 := len(nums2)

	if len1 > len2 {
		return findMedianSortedArrays(nums2, nums1)
	}

	k := (len1 + len2 + 1) / 2
	isOdd := false
	if (len1+len2)%2 == 1 {
		isOdd = true
	}

	l, r := 0, len1
	for l < r {
		m1 := l + (r-l)/2
		m2 := k - m1
		if nums2[m2-1] > nums1[m1] {
			l = m1 + 1
		} else {
			r = m1
		}
	}

	m1 := l
	m2 := k - l
	median := max(getValue(nums1, m1-1), getValue(nums2, m2-1))
	if isOdd {
		return float64(median)
	}
	median1 := min(getValue(nums1, m1), getValue(nums2, m2))
	return (float64(median1) + float64(median)) / 2
}

func getValue(nums []int, index int) int {
	if index < 0 {
		return INT_MIN
	}
	if index == len(nums) {
		return INT_MAX
	}
	return nums[index]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	nums1 := []int{1, 3, 4}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
