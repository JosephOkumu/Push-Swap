package lib

// Finds the index of the target value in a sorted array using binary search
func FindElementIndex(arr []int, targetValue int) int {
	start, end := 0, len(arr)-1
	for start <= end {
		mid := start + (end-start)>>1
		if arr[mid] == targetValue {
			return mid
		}
		if arr[mid] < targetValue {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
