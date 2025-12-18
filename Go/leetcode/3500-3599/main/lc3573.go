package main

// runtime: O(nk)
// space: O(k)
func maximumProfit(prices []int, k int) int64 {
	// dp[k][t]: 至多k次操作，t=0空，t=1hold，t=2sold
	// 最终k次操作空手是最佳情况
	dp := make([][3]int, k+1)
	// 第一天初始化操作
	for j := 1; j <= k; j++ {
		dp[j][1] = -prices[0]
		dp[j][2] = prices[0]
	}
	for _, p := range prices {
		for j := k; j > 0; j-- {
			dp[j][0] = max(dp[j][0], dp[j][1]+p, dp[j][2]-p)
			dp[j][1] = max(dp[j][1], dp[j-1][0]-p)
			dp[j][2] = max(dp[j][2], dp[j-1][0]+p)
		}
	}
	return int64(dp[k][0])
}
