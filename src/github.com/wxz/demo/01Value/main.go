package main

import "fmt"

// var 变量名 类型 = 表达式
// var name string = "Q1mi"
// var age int = 18

// var name string
// var num int
// var isOk bool

// 或者一次初始化多个变量
// var name, age = "Q1mi", 20

// 类型推导
// var name = "Q1mi"
// var age = 18

// var (
// 	a string
// 	b int
// 	c bool
// 	d float32
// )

// 全局变量m
var m = 100

func main() {
	n := 10
	m := 200 // 此处声明局部变量m
	fmt.Println(m, n)
}

// 匿名变量
// 在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。 匿名变量用一个下划线_表示，例如：

// func foo() (int, string) {
// 	return 10, "Q1mi"
// }
// func main() {
// 	x, _ := foo()
// 	_, y := foo()
// 	fmt.Println("x=", x)
// 	fmt.Println("y=", y)
// }

// 常量
// 相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。 常量的声明和变量声明非常类似，只是把var换成了const，常量在定义的时候必须赋值。

// const pi = 3.1415
// const e = 2.7182
// 声明了pi和e这两个常量之后，在整个程序运行期间它们的值都不能再发生变化了。

// 多个常量也可以一起声明：

// const (
//     pi = 3.1415
//     e = 2.7182
// )
// const同时声明多个常量时，如果省略了值则表示和上面一行的值相同。 例如：

// const (
//     n1 = 100
//     n2
//     n3
// )
// 上面示例中，常量n1、n2、n3的值都是100。

// iota
// iota是go语言的常量计数器，只能在常量的表达式中使用。

// iota在const关键字出现时将被重置为0。const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用。

// 举个例子：

// const (
// 		n1 = iota //0
// 		n2        //1
// 		n3        //2
// 		n4        //3
// 	)
// 几个常见的iota示例:
// 使用_跳过某些值

// const (
// 		n1 = iota //0
// 		n2        //1
// 		_
// 		n4        //3
// 	)
// iota声明中间插队

// const (
// 		n1 = iota //0
// 		n2 = 100  //100
// 		n3 = iota //2
// 		n4        //3
// 	)
// 	const n5 = iota //0
// 定义数量级 （这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）

// const (
// 		_  = iota
// 		KB = 1 << (10 * iota)
// 		MB = 1 << (10 * iota)
// 		GB = 1 << (10 * iota)
// 		TB = 1 << (10 * iota)
// 		PB = 1 << (10 * iota)
// 	)
// 多个iota定义在一行

// const (
// 		a, b = iota + 1, iota + 2 //1,2
// 		c, d                      //2,3
// 		e, f                      //3,4
// 	)
