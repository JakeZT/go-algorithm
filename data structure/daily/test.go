package main

import (
	"fmt"
)

func main() {
	//str := "abc"
	str := "hi,你好"

	//下标遍历
	strRune2 := []rune(str)
	for i, v := range strRune2 {
		fmt.Printf("str[%d]=%v\n", i, string(v))
	}

	//for...range遍历
	for i := 0; i < len(str); i++ {
		// fmt.Printf("str[%d]=%v\n", i, string(str[i]))
	}

	//转为[]rune类型，再下边遍历
	strRune := []rune(str)
	for i := 0; i < len(strRune); i++ {
		fmt.Printf("strRune[%d]=%v\n", i, string(strRune[i]))
	}

}
