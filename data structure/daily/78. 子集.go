package main

func main78() {}

// func subsets(nums []int) [][]int {
// 	res := [][]int{}
// 	var dfs func(i int, list []int)

// 	dfs = func(i int, list []int) {
// 		tmp := make([]int, len(list))
// 		copy(tmp, list)
// 		res = append(res, tmp)
// 		for j := i; j < len(nums); j++ {
// 			list = append(list, nums[j])
// 			dfs(j+1, list)
// 			list = list[:len(list)-1]
// 		}
// 	}

// 	dfs(0, []int{})
// 	return res
// }

// func subsets(nums []int) [][]int {
// 	res := &[][]int{}
// 	var dfs func(index int, list *[]int)
// 	dfs = func(index int, list *[]int) {
// 		if index == len(nums) {
// 			temp := append([]int{}, *list...)
// 			*res = append(*res, temp)
// 			return
// 		}
// 		*list = append(*list, nums[index])
// 		dfs(index+1, list)
// 		*list = (*list)[:len(*list)-1]
// 		dfs(index+1, list)
// 	}
// 	var cur []int
// 	dfs(0, &cur)
// 	return *res
// }

func subsets(nums []int) [][]int {
	res := [][]int{}
	var dfs func(index int, list []int)
	dfs = func(index int, list []int) {
		if index == len(nums) {
			temp := append([]int{}, list...)
			res = append(res, temp)
			return
		}
		list = append(list, nums[index])
		dfs(index+1, list)
		list = list[:len(list)-1]
		dfs(index+1, list)
	}
	dfs(0, []int{})
	return res
}
