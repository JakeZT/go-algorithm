package leetcode66

func main() {}

func plusOne(d []int) []int {
	for i := len(d) - 1; i >= 0; i-- {
		d[i]++
		d[i] = d[i] % 10
		if d[i] != 0 {
			return d
		}
	}
	d = make([]int, len(d)+1)
	d[0] = 1
	return d
}
