package main

import "fmt"

func main() {
	arr1 := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("### Merge Sort ###")
	margeSort(arr1, 0, len(arr1)-1)
	fmt.Printf("Sorted arr: %v\n", arr1)

	arr2 := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("### Quick Sort ###")
	quickSort(arr2, 0, len(arr2)-1)
	fmt.Printf("Sorted arr: %v\n", arr2)
}
