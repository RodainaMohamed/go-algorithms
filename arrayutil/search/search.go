package search

// LinearSearch returns an integer key index in an array.
func LinearSearch(key int, arr []int) int {
	for i, current := range arr {
		if current == key {
			return i
		}
	}
	return -1
}

// BinarySearch returns an integer key index in a sorted array.
func BinarySearch(key int, arr []int) int {
	var min, max, mid = 0, len(arr), 0
	for max >= min {
		mid = (max + min) / 2
		if arr[mid] == key {
			return mid
		}
		if arr[mid] < key {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	return -1
}
