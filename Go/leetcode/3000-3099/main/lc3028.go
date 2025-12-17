package main

// runtime: O(N)
// space: O(1)
func returnToBoundaryCount(nums []int) int {
	cnt, x := 0, 0
	for _, v := range nums {
		x += v
		if x == 0 {
			cnt++
		}
	}
	return cnt
}
