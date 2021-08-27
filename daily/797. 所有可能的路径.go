package main

import "fmt"

func main() {
	graph := [][]int{{1, 2}, {3}, {3}, {}}
	res := allPathsSourceTarget(graph)
	fmt.Println(res)
}

/* func allPathsSourceTarget(graph [][]int) [][]int {
	var (
		res     = [][]int{}
		n       = len(graph)
		visited = make([]bool, n)
	)

	var traverse func(int, []int)
	traverse = func(s int, path []int) {
		if s == n-1 {
			res = append(res, append([]int{}, path...))
			return
		}
		for _, j := range graph[s] {
			if !visited[j] {
				visited[j] = true
				path = append(path, j)
				traverse(j, path)
				path = path[:len(path)-1]
				visited[j] = false
			}
		}
	}
	visited[0] = true
	traverse(0, []int{0})
	return res
} */
func allPathsSourceTarget(graph [][]int) [][]int {
	var (
		res  = [][]int{}
		n    = len(graph)
		path = []int{}
	)
	var traverse func(int, []int)
	traverse = func(s int, path []int) {
		path = append(path, s)
		if s == n-1 {
			res = append(res, append([]int{}, path...))
			path = path[:len(path)-1]
			return
		}
		for _, node := range graph[s] {
			traverse(node, path)
		}
		path = path[:len(path)-1]
	}
	traverse(0, path)
	return res
}
