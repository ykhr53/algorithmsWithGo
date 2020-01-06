package main

import "fmt"

func main() {
	s1 := []int{2, 7, 9, 3, 1}
	fmt.Printf("### LC198 answer: %v \n", houseRobber(s1))

	s2 := []int{1, 1, 1, 1, 1}
	i2 := 3
	fmt.Printf("### LC494 answer: %v \n", findTargetSumWays(s2, i2))
}
