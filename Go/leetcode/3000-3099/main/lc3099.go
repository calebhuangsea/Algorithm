package main

// runtime: O(len(x))
// space: O(1)
func sumOfTheDigitsOfHarshadNumber(x int) int {
	sum := 0
	for temp := x; temp > 0; temp /= 10 {
		sum += temp % 10
	}
	if x%sum == 0 {
		return sum
	}
	return -1
}
