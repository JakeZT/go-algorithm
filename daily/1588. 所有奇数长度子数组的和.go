package leetcode1588

func main() {}

/*
给你一个正整数数组 arr ，请你计算所有可能的奇数长度子数组的和。

子数组 定义为原数组中的一个连续子序列。

请你返回 arr 中 所有奇数长度子数组的和 。

输入：arr = [1,4,2,5,3]
输出：58
解释：所有奇数长度子数组和它们的和为：
[1] = 1
[4] = 4
[2] = 2
[5] = 5
[3] = 3
[1,4,2] = 7
[4,2,5] = 11
[2,5,3] = 10
[1,4,2,5,3] = 15
我们将所有值求和得到 1 + 4 + 2 + 5 + 3 + 7 + 11 + 10 + 15 = 58

*/
func sumOddLengthSubarrays(arr []int) int {
	res, lens := 0, len(arr)
	for i := 1; i < lens; i++ {
		arr[i] += arr[i-1]
	}
	for i := 1; i <= lens; i += 2 {
		a, b := 0, i
		res += arr[b-1]
		for b < lens {
			res += arr[b] - arr[a]
			a++
			b++
		}
	}
	return res
}
