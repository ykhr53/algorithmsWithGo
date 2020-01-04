package main

import "fmt"

func merge(arr []int, left int, middle int, right int) {
	leftLength := middle - left + 1
	rightLength := right - middle

	leftArr := make([]int, leftLength)
	rightArr := make([]int, rightLength)

	for i := 0; i < leftLength; i++ {
		leftArr[i] = arr[left+i]
	}

	for j := 0; j < rightLength; j++ {
		rightArr[j] = arr[middle+1+j]
	}

	leftCounter := 0
	rightCounter := 0
	idx := left

	// filling from smaller element
	for leftCounter < leftLength && rightCounter < rightLength {
		if leftArr[leftCounter] <= rightArr[rightCounter] {
			arr[idx] = leftArr[leftCounter]
			leftCounter++
		} else {
			arr[idx] = rightArr[rightCounter]
			rightCounter++
		}
		idx++
	}

	// filling remaining elements
	for leftCounter < leftLength {
		arr[idx] = leftArr[leftCounter]
		leftCounter++
		idx++
	}
	for rightCounter < rightLength {
		arr[idx] = rightArr[rightCounter]
		rightCounter++
		idx++
	}

	fmt.Printf("merged result: %v\n", arr[left:right+1])
}

func margeSort(arr []int, left int, right int) {
	if left < right {
		middle := (right-left)/2 + left
		fmt.Printf("left side: %v\n", arr[left:middle+1])
		margeSort(arr, left, middle)
		fmt.Printf("right side: %v\n", arr[middle+1:right+1])
		margeSort(arr, middle+1, right)
		merge(arr, left, middle, right)
	}
}
