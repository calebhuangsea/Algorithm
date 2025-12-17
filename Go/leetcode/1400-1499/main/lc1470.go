package main

// runtime: O(N)
// space: O(N)
func shuffle(nums []int, n int) []int {
	res := make([]int, n*2)
	for i := 0; i < n; i++ {
		res[i*2], res[i*2+1] = nums[i], nums[n+i]
	}
	return res
}

func main() {
	println(shuffle([]int{2, 5, 1, 3, 4, 7}, 3))
}
