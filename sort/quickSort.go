package main

// QuickSort 通用快速排序
func QuickSort[T any](arr []T, compareFn func(a, b T) bool) {
	if len(arr) < 2 {
		return
	}

	pivot := arr[0]
	left := 1
	right := len(arr) - 1

	for left <= right {
		if compareFn(arr[left], pivot) {
			left++
		} else if compareFn(pivot, arr[right]) {
			right--
		} else {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}

	arr[0], arr[right] = arr[right], arr[0]

	QuickSort(arr[:right], compareFn)
	QuickSort(arr[right+1:], compareFn)
}
