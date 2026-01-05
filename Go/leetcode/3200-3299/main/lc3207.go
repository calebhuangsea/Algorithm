package main

import "math"

// runtime: O(N)
// space: O(1)
func maximumPoints(enemyEnergies []int, currentEnergy int) int64 {
	minE, sum := math.MaxInt, 0
	for _, v := range enemyEnergies {
		minE = min(minE, v)
		sum += v
	}
	if currentEnergy < minE {
		return 0
	}
	return int64((currentEnergy + sum - minE) / minE)
}
