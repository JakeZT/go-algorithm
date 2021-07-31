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

func arrSort(pq [][]int) {
	sort.SliceStable(pq, func(a, b int) bool {
		if pq[a][0] == pq[b][0] {
			return pq[a][1] < pq[b][1]
		} else {
			return pq[b][0] <= pq[b][0]
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
	arr = append(arr[:idx], arr[idx+1:]...)
	return arr
}

func myMax(a []int) int {
	m := 0
	for _, i := range a {
		if i > m {
			m = i
		}
	}
	return m
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
