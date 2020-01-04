package main

import "fmt"

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	margeSort(arr, 0, len(arr)-1)
	fmt.Printf("Sorted arr: %v\n", arr)
}
