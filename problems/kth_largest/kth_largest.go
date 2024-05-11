package kthlargest

import "golang.org/x/exp/slices"

func KthLargest(arr []int, k int) int {
	// arr := []int{4, 2, 9, 7, 5, 6, 7, 1, 3}
	// 1,2,3,4,5,6,7,7,9
	if k > len(arr) {
		return -1
	}
	k--

	// sort
	slices.Sort(arr)

	// return index k
	return arr[len(arr)-1-k]
}
