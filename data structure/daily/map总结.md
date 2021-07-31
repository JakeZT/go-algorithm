自定义 set 结构， 混合类型的数组 去重

map 如果用 for range 遍历, 顺序每次是随机的，return 或者 break 可以跳出 for range

原生方法没有 linkedHashMap，只能自己实现，或者用第三方包

```go
func mySet(arr []interface{}) (res []interface{}) {
	hashSet := make(map[interface{}]struct{})
	for _, v := range arr { //始终无序，随机访问， break或者return基本无用
		hashSet[v] = struct{}{}
	}
	for key := range hashSet {
		res = append(res, key)
	}
	fmt.Println(res) // [3213 Hello World 213]
	return
}
```
