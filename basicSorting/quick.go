package main

// In this code, I'm gonna always take tne last element as a pivot.
func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	i := left - 1

	for j := left; j < right; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[right] = arr[right], arr[i+1]
	return i + 1
}

func quickSort(arr []int, left int, right int) {
	if left < right {
		divideIdx := partition(arr, left, right)
		quickSort(arr, left, divideIdx-1)
		quickSort(arr, divideIdx+1, right)
	}
}
