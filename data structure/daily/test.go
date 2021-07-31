package main

import "fmt"

func mySet(arr []interface{}) (res []interface{}) {
	hashSet := make(map[interface{}]struct{})
	for _, v := range arr {
		hashSet[v] = struct{}{}
	}
	for key := range hashSet {
		res = append(res, key)
		if key == 111 {
			break
		}
	}
	return
}
func mainT() {
	data := []interface{}{"Hello", "World", 111, 213, 3213, 213, "World"}
	var res = mySet(data)
	fmt.Println(res)
}
