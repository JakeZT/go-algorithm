package main

import (
	"sort"
)

func main218() {
	temp := [][]int{
		{2, 9, 10}, {3, 7, 15}, {5, 12, 12}, {15, 20, 10}, {19, 24, 8},
	}
	// temp=[[2,9,10],[3,7,15],[5,12,12],[15,20,10],[19,24,8]]
	getSkyline(temp)
}

func myMax(arr []int) (_max int) {
	for _, i := range arr {
		if i > _max {
			_max = i
		}
	}
	return
}
func arrSort(pq [][]int) {
	sort.Slice(pq, func(i, j int) bool {
		if pq[i][0] == pq[j][0] {
			return pq[i][1] < pq[j][1]
		} else {
			return pq[i][0] <= pq[j][0]
		}
	})
}

func remove(arr []int, tar int) []int {
	idx := -1
	for key := 0; key <= len(arr)-1; key++ {
		if arr[key] == tar {
			idx = key
		}
	}
	if idx == -1 {
		return arr
	}
	arr = append(arr[:idx], arr[idx+1:]...)
	return arr
}

func getSkyline(buildings [][]int) [][]int {
	var (
		res = [][]int{}
		pq  = [][]int{}
		pre = -1
	)
	for _, b := range buildings {
		cur := [][]int{
			{b[0], -b[2]},
			{b[1], b[2]},
		}
		pq = append(pq, cur...)
	}
	arrSort(pq)
	heights := []int{0}
	for _, h := range pq {
		if h[1] < 0 {
			heights = append(heights, -h[1])
		} else {
			heights = remove(heights, h[1])
		}
		maxHeight := myMax(heights)
		if pre != maxHeight {
			cur := []int{h[0], maxHeight}
			res = append(res, cur)
			pre = maxHeight
		}
	}
	return res

}
