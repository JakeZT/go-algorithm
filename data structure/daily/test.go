package main

import (
	"fmt"

	"github.com/thoas/go-funk"
)

func main() {
	funk.IndexOf([]string{"foo", "bar"}, "bar") // 1
	var res int
	res = funk.IndexOf([]string{"foo", "bar"}, func(value string) bool {
		return value == "bar"
	}) // 1
	fmt.Println(res)
	funk.IndexOf([]string{"foo", "bar"}, "gilles") // -1
}
