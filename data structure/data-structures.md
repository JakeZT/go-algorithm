```go
for循环遍历数组
for i := 0; i < len(arr); i++ {
    //arrHaiCoder[i]
}

for index, value := range arrHaiCoder{
}
// 通过 _ 的形式忽略
for _, value := range arrHaiCoder{
}
```

遍历对象

```go
Go 语言 中 map 的遍历只能使用 for range 的形式
for key, value := range mapName{
}

遍历key
for key := range mapName{
}

遍历value
for _,value := range mapName{
}
```

for采用byte类型循环，for range采用rune类型循环









### 变量地址

* 变量本质就是内存中一块数据的标记.把值存储到变量中实质是把值存储到内存中
* 每次对变量重新赋值就是在修改变量地址中的内容
* 在Go语言中可以通过 `&+变量名` 获取到变量地址值
* 重新创建一个非引用型变量(即使是把已有变量直接赋值给新变量)也会新开辟内存地址.

```go
func main() {
	a := 3
	fmt.Println(&a) //输出:地址
	a = 4
	fmt.Println(&a) //输出的地址不变

	b := a
	b = 5
	fmt.Println(&b, &a) //两个值不相同
	fmt.Println(b, a)   //输出:5 4
}
```

### 指针变量

* 指针变量指向一个值的内存地址
* 使用&+变量 返回值就是一个指针类型
* 使用`var 变量名 *类型` 声明指针类型变量
* 声明指针不会开辟内存地址,只是准备要指向内存某个空间,而声明变量会开辟内存地址,准备存放内容.所以指针类型变量都是把一个变量的地址赋值给指针变量
* 使用`*+指针`能够获取内存地址中的值.所以`*+指针`就和直接使用变量是相同的
* 应用指针可以实现多个地方操作同一个内存地址的值(在方法参数应用指针较多)

```go
func main() {
	//创建变量
	a := 123
	var point *int
	point = &a //此时point和&a是相等的
	fmt.Println(point)
	*point = 3             //等价于a=3
	fmt.Println(*point, a) //*point和a是相当的
}
```

### 空指针

* 指针目的就是指向内存中一块地址
* 声明指针后指针不会指向任何内存地址,所以此时指针是空.在Go语言中空用nil表示

```go
func main() {
	var a *int
	fmt.Println(a)        //输出:<nil>
	fmt.Println(a == nil) //输出true
}
```



### new函数

* 在上一小节中学习了指针,每次创建一个指针必须在额外创建一个变量,操作比较麻烦.
* 可以通过new函数直接创建一个类型的指针

```
变量名:=new(Type)
```

* 使用new函数创建的指针已有指向,可以使用`*指针对象`进行赋值.

```go
func main() {
	a := new(int)
	fmt.Println(a) //输出:指针地址
	*a = 123
	fmt.Println(*a) //输出:123
}
```

* 只声明的指针变量不能直接赋值,

```go
func main() {
	var a *int
	*a = 123
	fmt.Println(*a)
}
```





### siwtch穿透和中断

* switch结构中某个最多只能执行一个case,使用fallthrough可以让下一个case/default继续执行

```go
func main() {
	switch num := 1; num {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("不是1,2,3,4")
	}
	fmt.Println("程序结束")
}
```

* break可以用在switch和循环中,表示立即结束,无论当前结构后面还有多少代码

```go
func main() {
	switch num := 1; num {
	case 1:
		fmt.Println("1")
		break
		fmt.Println("break后面代码都不执行")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("不是1,2,3,4")
	}
	fmt.Println("程序结束")
}
```

