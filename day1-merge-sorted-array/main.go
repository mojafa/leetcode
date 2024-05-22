package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	fmt.Println(nums1)
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

}

// merge merges two sorted arrays nums1 and nums2 into nums1 in-place.

func merge(nums1 []int, m int, nums2 []int, n int) {
	// i points to the last element in the initial segment of nums1
	i := m - 1
	// j points to the last element in nums2

	j := n - 1
	// k points to the last position in nums1

	k := m + n - 1
	// Merge nums1 and nums2 from the end

	for i >= 0 && j >= 0 {
		// Compare the last elements of nums1 and nums2 and place the larger one at the end of nums1

		if nums1[i] > nums2[j] {
			fmt.Println(nums1)
			nums1[k] = nums1[i]
			fmt.Println(nums1)
			i--
			fmt.Println(nums1)
		} else {
			nums1[k] = nums2[j]
			fmt.Println(nums1)
			j--
			fmt.Println(nums1)
		}
		k--
		fmt.Println(nums1)
	}
	// If there are remaining elements in nums2, copy them to nums1

	for j >= 0 {
		fmt.Println(nums1)
		nums1[k] = nums2[j]
		fmt.Println(nums1)
		j--
		fmt.Println(nums1)
		k--
		fmt.Println(nums1)
	}
}
