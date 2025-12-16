package main

// runtime: O(N)
// space: O(1)
func romanToInt(s string) int {
	values := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	n := len(s)
	res := 0
	for i := 0; i < n; i++ {
		v := values[s[i]]
		// 如果不是最后一个字符，且当前值小于下一个值 => 做减法
		if i+1 < n && v < values[s[i+1]] {
			res -= v
		} else {
			res += v
		}
	}
	return res
}
