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





### Golang的封装，继承，多态

#### 封装

* 封装主要体现在两个方面:封装数据、封装业务
* Go语言中通过首字母大小控制访问权限.属性首字母小写对外提供访问方法是封装数据最常见的实现方式
* 可以通过方法封装业务
  * 提出方法是封装
  * 控制结构体属性访问,对外提供访问方法也是封装
* 在面向对象中封装的好处:
  * 安全性.结构体属性访问受到限制,必须按照特定访问渠道
  * 可复用性,封装的方法实现可复用性
  * 可读写,多段增加代码可读性

```go
type People struct { // 大写结构体名称，对外包可见
	name string //大写属性名，外包可以访问，小写则不可以访问
	age  int    //体重.单位斤
}

func (p *People) SetName(name string) {
	p.name = name
}
func (p *People) GetName() string {
	return p.name
}
```



### 结构体

+ 由于结构体是值类型,在方法传递时希望传递结构体地址,可以使用时结构体指针完成
+ **可以结合new(T)函数创建结构体指针**

```go
	peo := new(People)
	//因为结构体本质是值类型,所以创建结构体指针时已经开辟了内存空间
	fmt.Println(peo == nil) //输出:false
	//由于结构体中属性并不是指针类型,所以可以直接调用
	peo.Name = "smallming"
	fmt.Println(peo)//输出:&{smallming 0}
	peo1:=peo
	peo1.Name="佳明哥"
	fmt.Println(peo1,peo)//输出:&{佳明哥 0} &{佳明哥 0}
```

* 如果不想使用new(T)函数,可以直接声明结构体指针并赋值

* 结构体可以定义在函数内部或函数外部(与普通变量一样),定义位置影响到结构体的访问范围
* 如果结构体定义在函数外面,结构体名称首字母是否大写影响到结构体是否能跨包访问
* 如果结构体能跨包访问,属性首字母是否大写影响到属性是否跨包访问

```go
type People struct {
	Name string
	Age  int
}
```

* 声明结构体变量
  * 由于结构体是值类型,所以声明后就会开辟内存空间
  * 所有成员为类型对应的初始值
* 双等(==)判断结构体中内容是否相等，地址肯定不同，因为是值类型



* 结构体解释:将一个或多个变量组合到一起,形成新的类型.这个类型就是结构体
* Go语言中的结构体和C++结构体有点类似,而Java或C#中类本质就是结构体
* 结构体是值类型
* 结构体定义语法
  * 通过语法可以看出,Go语言发明者明确认为结构体就是一种自定义类型

```go
    type 结构体名称 struct{
      名称 类型//成员或属性
    }
```





### 方法的使用

```go
type People struct {
	Name string//姓名
	Weight	float64//体重.单位斤
}

func (p *People) run(){
	fmt.Println(p.Name,"正在跑步,体重为:",p.Weight)//输出:张三 正在跑步,体重为: 17
	p.Weight-=0.1
}

func main() {
	peo:=&People{"张三",17}
	peo.run()
	fmt.Println(peo.Name,"跑完步后的体重是",peo.Weight)//输出:张三 跑完步后的体重是 16.9
}
```



### 值传递和引用传递

* 讨论值传递和引用传递时,其实就是看值类型变量和引用类型变量作为函数参数时,修改形参是否会影响到实参
* **结构体是值类型**
* **在Go语言中五个引用类型变量,其他都是值类型**
  * slice
  * map
  * channel
  * interface
  * func()
* 引用类型作为参数时,称为浅拷贝,形参改变,实参数跟随变化.因为传递的是地址,形参和实参都指向同一块地址
* 值类型作为参数时,称为深拷贝,形参改变,实参不变,因为传递的是值的副本,形参会新开辟一块空间,与实参指向不同
* **如果希望值类型数据在修改形参时实参跟随变化,可以把参数设置为指针类型**

### 二.代码演示

* 值类型作为参数代码演示

```go
package main

import "fmt"

func demo(i int, s string) {
	i = 5
	s = "改变"
}

func main() {
	i := 1
	s := "原值"
	demo(i, s)
	fmt.Println(i, s) //输出:1 原值
}
```

* 引用传递代码示例

```go
package main

import "fmt"

func demo(arg []int) {
   arg[len(arg)-1] = 110
}

func main() {
   s := []int{1, 2, 3}
   demo(s)
   fmt.Println(s) //输出:[1 2 110]
}
```

* 如果希望值类型实参跟随形参变化,可以把值类型指针作为参数

```go
package main

import "fmt"

//行参指针类型
func demo(i *int, s string) {
   //需要在变量前面带有*表示指针变量
   *i = 5
   s = "改变"
}

func main() {
   i := 1
   s := "原值"
   //注意此处第一个参数是i的地址,前面&
   //s保留为值类型
   demo(&i, s)
   fmt.Println(i, s) //输出:5 原值
}
```





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



#### 跳转continue， break

* Go语言执行标签写法,可以通过定义标签让continue控制影响哪个for循环

```go
	myfor:for k := 0; k < 2; k++ {
		for i := 0; i < 3; i++ {
			if i == 1 {
				continue myfor
			}
			fmt.Println(k, i, "结束")
		}
	}
```

* break也可以通过定义标签,控制break对哪个for循环生效

```go
	myfor:for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if j == 1 {
				break myfor
			}
			fmt.Println(i, j)
		}
	}
```

### goto关键字

* goto是Go语言中的一个关键字
* goto让编译器执行时跳转到特定位置
  * Loop是标记名(Label),名称任意,习惯上名称为Loop

```go
	fmt.Println("执行程序")
	i := 6
	if i == 6 {
		goto Loop
	}
	fmt.Println("if下面输出")
Loop:
	fmt.Println("loop")
```

* 可以有多个,但是标签(Labal)定义了就必须使用

```go
	fmt.Println("执行程序")
	i := 6
	if i == 6 {
		goto Loop
	}
	fmt.Println("if下面输出")
Loop:
	fmt.Println("loop")
Loop1: //报错:label Loop1 defined and not used
	fmt.Println("Loop1")
```

* goto也可以用于跳出循环,执行指定标签位置代码

```go
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		if i == 2 {
			goto abc
		}
	}
	fmt.Println("for循环执行结束")
abc:
	fmt.Println("abc")
	fmt.Println("程序结束")
```



如果添加一次添加多个值,且添加后的长度大于扩容一次的大小,容量和长度相等.等到下次添加内容时如果不超出扩容大小,在现在的基础上进行翻倍

```go
s := make([]string, 0,9)
s = append(s,"10")
fmt.Println(len(s), cap(s)) //输出:10 18	s = append(s,"10")
fmt.Println(len(s), cap(s)) //输出:10 18
```



切片融合

```go
s := make([]string, 0)
s1 := []string{"smallming", "jake"}
s = append(s, s1...) //注意此处,必须有三个点
fmt.Println(s)
```

定义数组后,取出数组中一个片段,这个片段就是切片类型

切片是指针,指向数组元素地址,修改切片的内容,数组的内容会跟随变化

当切片内容在增加时

* 如果增加后切片的长度没有超出数组,修改切片也是在修改数组
* 如果增加后切片的长度超出数组,会重新开辟一块空间放切片的内容
* 通过下面代码也正面了切片中内容存在一块连续空间(和数组一样)





go中删除数组

```go

	num := []int {0,1,2,3,4,5,6}
	//要删除脚标为n的元素
	n:= 2
	num1 :=num[0:n]
	num1= append(num1,num[n+1:]...)
	fmt.Println(num1)
```

数组删除一个元素

```go
a := []int{0, 1, 2, 3, 4}
//删除第i个元素
i := 2
a = append(a[:i], a[i+1:]...)
```



### 使用copy完成删除元素

* 使用copy函数可以保证原切片内容不变

```go
	s := []int{1, 2, 3, 4, 5, 6, 7}
	n := 2 //要删除元素的索引
	newSlice := make([]int, n)
	copy(newSlice, s[0:n])
	newSlice = append(newSlice, s[n+1:]...)
	fmt.Println(s)        //原切片不变
	fmt.Println(newSlice) //删除指定元素后的切片
```







### 排序实现

* 对int类型切片排序

```go
	num := [] int{1, 7, 5, 2, 6}
	sort.Ints(num) //升序
	fmt.Println(num)
	sort.Sort(sort.Reverse(sort.IntSlice(num))) //降序
	fmt.Println(num)
```

* 对float64类型切片排序

```go
	f := [] float64{1.5, 7.2, 5.8, 2.3, 6.9}
	sort.Float64s(f) //升序
	fmt.Println(f)
	sort.Sort(sort.Reverse(sort.Float64Slice(f))) //降序
	fmt.Println(f)
```

* 对string类型切片排序
  * 按照编码表数值进行排序
  * 多字符串中按照第一个字符进行比较
  * 如果第一个字符相同,比较第二个字符

```go
	s := []string{"我", "我是中国人", "a", "d", "国家", "你", "我a"}
	sort.Sort(sort.StringSlice(s)) //升序
	fmt.Println(s)
	//查找内容的索引,如果不存在,返回内容应该在升序排序切片的哪个位置插入
	fmt.Println(sort.SearchStrings(s, "你是"))
	sort.Sort(sort.Reverse(sort.StringSlice(s)))//降序
	fmt.Println(s)
```



map中不存在的值，访问会返回空字符串+ false

```go
	m := map[string]string{"name": "smallming", "address": "北京海淀"}
	fmt.Println(m["name"]) //输出:smallming
	fmt.Println(m["age"])  //输出:空字符串
	value, ok := m["age"]
	fmt.Println(value, ok) //输出:空字符串 false
```



### 操作List

* 直接使用container/list包下的New()新建一个空的List

```go
	mylist := list.New()
	fmt.Println(mylist)       //输出list中内容
	fmt.Println(mylist.Len()) //查看链表中元素的个数
	fmt.Printf("%p", mylist)  //输出地址
```

* Go语言标准库中提供了很多向双向链表中添加元素的函数

```go
	//添加到最后,List["a"]
	mylist.PushBack("a")
    //添加到最前面,List["b","a"]
	mylist.PushFront("b") 
	//向第一个元素后面添加元素,List["b","c","a"]
	mylist.InsertAfter("c", mylist.Front()) 
	//向最后一个元素前面添加元素,List["b","c","d","a"]
	mylist.InsertBefore("d", mylist.Back()) 
```

* 取出链表中的元素

```go
	fmt.Println(mylist.Back().Value)  //最后一个元素的值
	fmt.Println(mylist.Front().Value) //第一个元素的值

	//只能从头向后找,或从后往前找,获取元素内容
	n := 5
	var curr *list.Element
	if n > 0 && n <= mylist.Len() {
		if n == 1 {
			curr = mylist.Front()
		} else if n == mylist.Len() {
			curr = mylist.Back()
		} else {
			curr = mylist.Front()
			for i := 1; i < n; i++ {
				curr = curr.Next()
			}
		}
	} else {
		fmt.Println("n的数值不对")
	}
	//遍历所有值
	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
```

* 移动元素的顺序

```go
	mylist.MoveToBack(mylist.Front()) //把第一个移动到后面
	mylist.MoveToFront(mylist.Back()) //把最后一个移动到前面
	mylist.MoveAfter(mylist.Front(),mylist.Back())//把第一个参数元素,移动到第二个参数元素后面
	mylist.MoveBefore(mylist.Front(),mylist.Back())//把第一个参数元素,移动到第二个参数元素前面
```

* 删除元素

```go
mylist.Remove(mylist.Front())
```

### 多返回值函数

* 在Go语言中每个函数声明时都可以定义成多返回值函数
* Go语言中所有的错误都是通过返回值返回的
* 声明多返回值函数的语法

```go
func 函数名(参数列表) (返回值,返回值){
  //函数体
}
```

* 调用函数的语法

```go
变量,变量:=函数名(参数)
```

* 调用函数时如果不想接收可以使用下划线占位

```go
变量,_:=函数名(参数)
```

* 理论上函数返回值个数可以无限多个,但是一般不去定义特别多个返回值(用结构体代替多返回值)







# 函数作为参数或返回值

* 变量可以作为函数的参数或返回值类型.而函数既然可以当做变量看待,函数变量也可以当做函数的参数或返回值
* 函数作为参数时,类型写成对应的类型即可

```go
func main() {
	a(func(s string) {
		fmt.Println(s)
	})
}

func a(b func(s string)) {
	fmt.Println("a执行")
	b("传递给s的内容")
}
```

* 函数作为返回值

```go
func main() {
	//此时result指向返回值函数.
	result := a()
	//调用函数,才能获取结果
	fmt.Println(result())
}

func a() func() int {
	return func() int {
		return 110
	}
}
```

### 包概述

* 包(package)是Go语言中组织单元.包是逻辑上的分组.而物理上的分组是不同的文件夹,文件夹和包一般是对应的
* 把多个文件放入同一个文件夹中,这些文件就是在同一个包中.
* 虽然允许源码文件的package和文件夹名不同但是最终编译后都会把文件的package编译成文件夹名称.所以为防止错误最好把文件的package和文件夹名称设置成相同的
* 一个Go语言项目必须要有main包,其他自定义名称的包个数任意,根据自己的需求即可.
* Go语言在寻找包时会从GOPATH/src 路径中寻找包,如果不存在去GOROOT/src(Go语言标准库源码所在文件夹)下找
* 不同包下资源可以相互访问,在导入其他包后,可以访问包下首字母大写的内容
* 同包下不同文件中全局资源可以随意访问

#### 自定义包

* 新建项目后在项目下新建src文件夹,在src文件夹中新建demo文件
* 在demo文件中新建demo1.go和demo2.go文件
* demo1.go文件源码如下

```go
package demo//包为demo

import "fmt"

func demo1(){
	fmt.Println("执行demo1")
}
```

* demo2.go文件源码如下

```go
package demo//包为demo

import "fmt"

func Demo2()  {//函数名大写才能被其他包访问
	fmt.Println("执行demo2")
	demo1()//同包下内容任意访问
}
```

* 在项目根目录下新建main.go,源码如下

```go
package main

import "demo"

func main() {
	demo.Demo2()
}
```

* 运行整个项目后,发现可以正常调用Demo2()函数
* 整个程序目录结构如下





### 全局变量

* 全局变量声明到函数外部,整个包都可以访问
* 如果全局变量首字母大写,跨包也可以访问.
* 声明全局变量时规范是

```go
var (
	变量名
	变量名=值
)
```

* 全局变量代码示例

```go
var (
	name = "smallming"
	age  = 17
)

func demo1() {
  	fmt.Println("名字:",name)
}

func demo2() {
  	fmt.Println("年龄:",age)
}
```

### 闭包概述

* 闭包不是Go语言独有的概念,在很多编程语言中都有闭包
* 闭包就是解决局部变量不能被外部访问一种解决方案
* 是把函数当作返回值的一种应用

#### 代码演示

* 总体思想为:在函数内部定义局部变量,把另一个函数当作返回值,局部变量对于返回值函数就相当于全局变量,所以多次调用返回值函数局部变量的值跟随变化

```go
package main

import "fmt"

func main() {
	//res其实就是test1返回值函数,和之前匿名函数变量一个道理
	res := test1()
	fmt.Println(res()) //输出2
	fmt.Println(res()) //输出3
	fmt.Println(res()) //输出4
}

//注意此处,返回值类型是func int
func test1() func() int {
	i := 1
	return func() int {
		i = i + 1
		return i
	}
}
```

* 如果重新调用test1()会重新声明及赋值局部变量i

```go
package main

import "fmt"

func main() {
	f := test1()
	fmt.Println("f的地址", f) //输出匿名函数地址
	fmt.Println("f:", f()) //调用匿名函数输出2
	fmt.Println("f:", f()) //调用匿名函数输出3
	k := test1()
	fmt.Println("k的地址", k) //输出匿名函数地址,与f相等
	fmt.Println("k:", k()) //调用匿名函数输出2
	fmt.Println("f:", f()) //输出:4
	fmt.Println("k:", k()) //输出:3
}

func test1() func() int {
	i := 1
	return func() int {
		i++
		// 每调用一次test1()输出的地址不一样
		fmt.Println("i的地址:", &i)
		return i
	}
}

```





