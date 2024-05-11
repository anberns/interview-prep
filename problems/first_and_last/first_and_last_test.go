package firstandlast_test

import (
	"testing"

	firstandlast "example.com/first_and_last"
	"github.com/stretchr/testify/require"
)

func TestFirstAndLast(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 5, 5, 5, 7, 8, 9}

	// target exists
	target := 5
	result := firstandlast.FirstAndLast(arr, target)
	require.Equal(t, []int{4, 7}, result)

	// target exist, single element
	target = 4
	result = firstandlast.FirstAndLast(arr, target)
	require.Equal(t, []int{3, 3}, result)

	// target doesn't exist
	target = 6
	result = firstandlast.FirstAndLast(arr, target)
	require.Equal(t, []int{-1, -1}, result)

	// empty array
	emptyArr := []int{}
	target = 6
	result = firstandlast.FirstAndLast(emptyArr, target)
	require.Equal(t, []int{-1, -1}, result)
}

func TestFirstAndLastBinary(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 5, 5, 5, 7, 8, 9}

	// target exists
	target := 5
	result := firstandlast.FirstAndLastBinary(arr, target)
	require.Equal(t, []int{4, 7}, result)

	// target exist, single element
	target = 4
	result = firstandlast.FirstAndLastBinary(arr, target)
	require.Equal(t, []int{3, 3}, result)

	// target doesn't exist
	target = 6
	result = firstandlast.FirstAndLastBinary(arr, target)
	require.Equal(t, []int{-1, -1}, result)

	// empty array
	emptyArr := []int{}
	target = 6
	result = firstandlast.FirstAndLastBinary(emptyArr, target)
	require.Equal(t, []int{-1, -1}, result)
}
