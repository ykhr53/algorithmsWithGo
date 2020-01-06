package main

func findTargetSumWays(nums []int, S int) int {
	dp := make(map[int]int)
	dp[0] = 1
	for _, num := range nums {
		m := make(map[int]int)
		for k := range dp {
			num1 := k + num
			m[num1] = getOrDefault(m, num1, 0) + dp[k]
			num2 := k - num
			m[num2] = getOrDefault(m, num2, 0) + dp[k]
		}
		dp = m
	}
	return getOrDefault(dp, S, 0)
}

func getOrDefault(m map[int]int, key int, def int) int {
	if v, ok := m[key]; ok {
		return v
	} else {
		return def
	}
}
