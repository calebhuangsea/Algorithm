package main

// runtime: O(N)
// space: O(1)
func getDescentPeriods(prices []int) int64 {
	var res, cnt int64 = 0, 0
	for i := 0; i < len(prices); i++ {
		if i == 0 || prices[i]+1 == prices[i-1] {
			cnt++
		} else {
			cnt = 1
		}
		res += cnt
	}
	return res
}
