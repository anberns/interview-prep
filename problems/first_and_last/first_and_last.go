package firstandlast

func FirstAndLast(arr []int, target int) []int {

	first := -1
	last := -1

	firstFound := false
	lastFound := false

	for i := 0; i < len(arr); i++ {
		if arr[i] == target {
			if !firstFound {
				first = i
				firstFound = true
			}

			last = i
			lastFound = true
		} else {
			if lastFound {
				break
			}
		}
	}

	out := []int{first, last}

	return out
}

func FirstAndLastBinary(arr []int, target int) []int {
	out := []int{-1, -1}

	bsFirst(arr, out, 0, len(arr)-1, target)
	bsLast(arr, out, 0, len(arr)-1, target)

	return out
}

func bsFirst(arr, out []int, start, finish, val int) {
	// base case
	if start > finish {
		return
	}

	// get midpoint (starting index plus halfway between start and finish)
	mid := start + ((finish - start) / 2)

	// if value is there, return index
	if arr[mid] == val && mid > 0 && arr[mid-1] != val {
		out[0] = mid
	}

	// if value is less than mid, call on lower half
	if arr[mid] > val || (arr[mid] == val && mid > 0 && arr[mid-1] == val) {
		bsFirst(arr, out, start, mid-1, val)
	} else {
		bsFirst(arr, out, mid+1, finish, val)
	}

}

func bsLast(arr, out []int, start, finish, val int) {
	// base case
	if start > finish {
		return
	}

	// get midpoint (starting index plus halfway between start and finish)
	mid := start + ((finish - start) / 2)

	if arr[mid] == val && mid+1 <= len(arr)-1 && arr[mid+1] != val {
		out[1] = mid
	}

	// if value is less than mid, call on lower half
	if arr[mid] < val || (arr[mid] == val && mid+1 <= len(arr)-1 && arr[mid+1] == val) {
		bsLast(arr, out, mid+1, finish, val)
	} else {
		bsLast(arr, out, start, mid-1, val)
	}

}
