package main

import "strings"

// runtime: O(N)
// space: O(N)
func defangIPaddr(address string) string {
	res := strings.Builder{}
	res.Grow(len(address) + 8)
	for _, s := range address {
		if s == '.' {
			res.WriteString("[.]")
		} else {
			res.WriteByte(byte(s))
		}
	}
	return res.String()
}

func main() {
	println(defangIPaddr("1.1.1.1"))
	println(defangIPaddr("255.100.50.0"))
}
