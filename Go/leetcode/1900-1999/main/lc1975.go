package main

import "math"

// runtime: O(N*M)
// space: O(1)
func maxMatrixSum(matrix [][]int) int64 {
	minX, count, res := math.MaxInt, 0, 0
	for _, ma := range matrix {
		for _, v := range ma {
			if v < 0 {
				count++
				v = -v
			}
			res += v
			minX = min(minX, v)
		}
	}
	if count%2 == 1 {
		res -= 2 * minX
	}
	return int64(res)
}
