package binarysearch

// accepts a sorted array and a value and returns the index if found or -1
func BinarySearch(arr []int, val int) int {

	// [1,2,4,6,7,8,10,14]

	return bSearch(arr, 0, len(arr)-1, val)
}

func bSearch(arr []int, start, finish, val int) int {
	// base case
	if start > finish {
		return -1
	}

	// get midpoint (starting index plus halfway between start and finish)
	mid := start + ((finish - start) / 2)

	// if value is there, return index
	if arr[mid] == val {
		return mid
	}

	// if value is less than mid, call on lower half
	if arr[mid] > val {
		return bSearch(arr, start, mid-1, val)
	} else {
		return bSearch(arr, mid+1, finish, val)
	}

}
