package main

// runtime: O(N)
// space: O(1)
func lengthOfLongestSubstring(s string) (ans int) {
	cnt := [128]int{}
	left := 0
	for right, c := range s {
		cnt[c]++
		for cnt[c] > 1 {
			cnt[s[left]]--
			left++
		}
		ans = max(ans, right-left+1)
	}
	return
}
