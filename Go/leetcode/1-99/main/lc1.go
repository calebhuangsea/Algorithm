package main

// runtime：O(N)
// space: O(N)
func twoSum(nums []int, target int) []int {
	cacheMap := make(map[int]int)
	for i, v := range nums {
		if j, ok := cacheMap[target-v]; ok {
			return []int{j, i}
		}
		cacheMap[v] = i
	}
	return []int{-1, -1}
}
