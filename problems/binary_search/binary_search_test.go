package binarysearch_test

import (
	"testing"

	binarysearch "example.com/binary_search"
	"github.com/stretchr/testify/require"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 2, 4, 5, 7, 8, 10, 12}

	// first element
	val := 1
	result := binarysearch.BinarySearch(arr, val)
	require.Equal(t, 0, result)

	// last element
	val = 12
	result = binarysearch.BinarySearch(arr, val)
	require.Equal(t, 7, result)

	// middle element
	val = 5
	result = binarysearch.BinarySearch(arr, val)
	require.Equal(t, 3, result)

	// element not in array lower
	val = -1
	result = binarysearch.BinarySearch(arr, val)
	require.Equal(t, -1, result)

	// element not in array upper
	val = 14
	result = binarysearch.BinarySearch(arr, val)
	require.Equal(t, -1, result)

	// element not in array middle
	val = 6
	result = binarysearch.BinarySearch(arr, val)
	require.Equal(t, -1, result)

	// zero length array
	emptyArr := []int{}
	val = 6
	result = binarysearch.BinarySearch(emptyArr, val)
	require.Equal(t, -1, result)

	// single length array, contains
	singleArr := []int{6}
	val = 6
	result = binarysearch.BinarySearch(singleArr, val)
	require.Equal(t, 0, result)

	// single length array, doesn't contain
	singleArr = []int{5}
	val = 6
	result = binarysearch.BinarySearch(singleArr, val)
	require.Equal(t, -1, result)
}
