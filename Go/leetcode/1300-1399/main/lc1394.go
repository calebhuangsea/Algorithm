package main

// runtime: O(N)
// space: O(N)
func findLucky(arr []int) int {
	cnt := make(map[int]int, len(arr))
	for _, v := range arr {
		cnt[v]++
	}
	res := -1
	for k, v := range cnt {
		if k == v && res < k {
			res = k
		}
	}
	return res
}
