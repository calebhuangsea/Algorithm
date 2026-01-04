package main

const max = 1e5 + 1

var cnt, sum [max]int

func init() {
	for i := 1; i < max; i++ {
		for j := i; j < max; j += i {
			cnt[j]++
			sum[j] += i
		}
	}
}

// runtime: O(N + max*log(max))
// space: O(max)
func sumFourDivisors(nums []int) int {
	res := 0
	for _, v := range nums {
		if cnt[v] == 4 {
			res += sum[v]
		}
	}
	return res
}
