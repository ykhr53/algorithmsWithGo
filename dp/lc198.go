package main

func houseRobber(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		}
		return nums[1]
	}
	N := len(nums)
	dp := make([]int, N)

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < N; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[N-1]
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
