package searching

// BinarySearch returns the index of el in a sorted (non-decreasing) slice arr.
//
// If the element is not present, the function returns -1.
//
// Requirements:
//   - arr must be sorted in non-decreasing order
//   - time complexity: O(log n)
//   - space complexity: O(1)
//
// Examples:
//
//	searching.BinarySearch([]int{1, 3, 5, 7, 9}, 5) // returns 2
//	searching.BinarySearch([]int{1, 2, 3}, 4)        // returns -1
//	searching.BinarySearch([]int{}, 10)              // returns -1
//
// When multiple identical elements exist, any of their indices may be returned.
// The search is performed using the classic iterative way without recursion.
func BinarySearch(arr []int, el int) int {
	low := 0
	high := len(arr) - 1
	var mid int

	for low <= high {
		mid = (low + high) / 2 // midpoint calculation

		if arr[mid] == el {
			return mid
		} else if arr[mid] > el {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1 // not found
}
