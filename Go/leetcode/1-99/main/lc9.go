package main

//import "strconv"
//// runtime: O(log_10(x))
//// space: O(log_10(x))
//func isPalindrome(x int) bool {
//	s := strconv.Itoa(x)
//	for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
//		if s[l] != s[r] {
//			return false
//		}
//	}
//	return true
//}

// runtime: O(log_10(x))
// space: O(log_10(x))
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	reverseX := 0
	for temp := x; temp > 0; temp /= 10 {
		reverseX = reverseX*10 + temp%10
	}
	return reverseX == x
}
