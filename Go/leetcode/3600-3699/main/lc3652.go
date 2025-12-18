package main

// runtime: O(N)
// space: O(1)
func maxProfit(prices []int, strategy []int, k int) int64 {
	// 定长滑窗，初始化第一个K窗口，每次往右移的时候，右边转为卖出，左边恢复
	n := len(prices)
	interval, sum := 0, 0
	for i, v := range prices {
		sum += v * strategy[i]
	}
	res := sum
	for i := 0; i < k; i++ {
		if i >= k/2 {
			interval += prices[i]
		}
		sum -= prices[i] * strategy[i]
	}
	res = max(res, sum+interval)
	for i := k; i < n; i++ {
		sum -= prices[i] * strategy[i]
		interval = interval + prices[i] - prices[i-k/2] + strategy[i-k]*prices[i-k]
		res = max(res, sum+interval)
	}
	return int64(res)
}

func main() {
	// 10
	println(maxProfit([]int{4, 7, 13}, []int{-1, -1, 0}, 2))
}
