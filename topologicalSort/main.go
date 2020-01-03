package main

import "fmt"

func main() {
	numCourses := 4
	prerequisites := [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}
	lc210ans := findOrder(numCourses, prerequisites)
	fmt.Printf("LC 210 answer: %v\n", lc210ans)
}
