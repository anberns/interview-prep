package kthlargest_test

import (
	"testing"

	kthlargest "example.com/kth_largest"
	"github.com/stretchr/testify/require"
)

func TestKthLargest(t *testing.T) {
	arr := []int{4, 2, 9, 7, 5, 6, 7, 1, 3}
	k := 4

	// middle
	result := kthlargest.KthLargest(arr, k)
	require.Equal(t, 6, result)

	// largest
	k = 1
	result = kthlargest.KthLargest(arr, k)
	require.Equal(t, 9, result)

	// smallest
	k = 9
	result = kthlargest.KthLargest(arr, k)
	require.Equal(t, 1, result)

	// k out of bounds
	k = 10
	result = kthlargest.KthLargest(arr, k)
	require.Equal(t, -1, result)
}
