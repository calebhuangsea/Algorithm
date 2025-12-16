package main

// runtime: O(N)
// space: O(1)
func minOperations(nums []int, k int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum % k
}
