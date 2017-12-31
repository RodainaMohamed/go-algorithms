package sort

// BubbleSort sorts the integer array ascendingly using bubble sort algorithm.
func BubbleSort(arr []int) {
	var tmp int
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i; j < length-i-1; j++ {
			if arr[j+1] < arr[j] {
				tmp = arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = tmp
			}
		}
	}
}

// SelectionSort sorts the integer array ascendingly using selection sort algorithm.
func SelectionSort(arr []int) {
	var tmp, min int
	length := len(arr)
	for i := 0; i < length-1; i++ {
		min = i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			tmp = arr[i]
			arr[i] = arr[min]
			arr[min] = tmp
		}
	}
}

// DoubleSelectionSort sorts the integer array ascendingly using double selection sort algorithm.
func DoubleSelectionSort(arr []int) {
	var min, max, oldMax, tmp int
	length := len(arr)
	for i := 0; i < length/2; i++ {
		min = i
		oldMax = length - i - 1
		max = oldMax
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[min] {
				min = j
			}
			if arr[length-j] > arr[max] && (min != max) {
				max = length - j
			}
		}
		if min != i {
			tmp = arr[i]
			arr[i] = arr[min]
			arr[min] = tmp
		}
		if max != oldMax {
			tmp = arr[oldMax]
			arr[oldMax] = arr[max]
			arr[max] = tmp
		}
	}
}

// InsertionSort sorts the integer array ascendingly using insertion sort algorithm.
func InsertionSort(arr []int) {
	var tmp int
	length := len(arr)
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if arr[j] < arr[j-1] {
				tmp = arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = tmp
			}
		}
	}
}
