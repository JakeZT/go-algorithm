package leetcode67

import (
	"math/big"
	"strconv"
)

func main() {}

/*
输入: a = "1010", b = "1011"
输出: "10101"
*/

func addBinary(a string, b string) string {
	ans := ""
	carry := 0
	n, lenA, lenB := len(a), len(a), len(b)
	if lenB > n {
		n = lenB
	}

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}
func addBinary1(a string, b string) string {

	ai, _ := new(big.Int).SetString(a, 2)
	bi, _ := new(big.Int).SetString(b, 2)

	ai.Add(ai, bi)
	return ai.Text(2)
}
