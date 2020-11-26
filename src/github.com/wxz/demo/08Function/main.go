package main

import "fmt"

func main() {
	// sayHello()
	// ret := intSum(10, 20)
	// fmt.Println(ret)

	// ret1 := intSum2()
	// ret2 := intSum2(10)
	// ret3 := intSum2(10, 20)
	// ret4 := intSum2(10, 20, 30)
	// fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60

	// ret5 := intSum3(100)
	// ret6 := intSum3(100, 10)
	// ret7 := intSum3(100, 10, 20)
	// ret8 := intSum3(100, 10, 20, 30)
	// fmt.Println(ret5, ret6, ret7, ret8) //100 110 130 160

	// testGlobalVar()

	// testLocalVar1()
	// fmt.Println(x) // 此时无法使用变量x

	// testNum()
	// fmt.Println(num)

	// testLocalVar3()

	// var c calculation               // 声明一个calculation类型的变量c
	// c = add                         // 把add赋值给c
	// fmt.Printf("type of c:%T\n", c) // type of c:main.calculation
	// fmt.Println(c(1, 2))            // 像调用add一样调用c

	// f := add                        // 将函数add赋值给变量f1
	// fmt.Printf("type of f:%T\n", f) // type of f:func(int, int) int
	// fmt.Println(f(10, 20))          // 像调用add一样调用f

	ret2 := calc(10, 20, add)
	fmt.Println(ret2) //30
}

// func sayHello() {
// 	fmt.Println("hello")
// }

// func intSum(x int, y int) int { // 可以写成"x , y int"类型简写
// 	return x + y
// }

// 可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。
// 注意：可变参数通常要作为函数的最后一个参数。
// func intSum2(x ...int) int {
// 	fmt.Println(x) //x是一个切片
// 	sum := 0
// 	for _, v := range x {
// 		sum = sum + v
// 	}
// 	return sum
// }

// 固定参数搭配可变参数使用时，可变参数要放在固定参数的后面
// func intSum3(x int, y ...int) int {
// 	fmt.Println(x, y)
// 	sum := x
// 	for _, v := range y {
// 		sum = sum + v
// 	}
// 	return sum
// }

// 本质上，函数的可变参数是通过切片来实现的。

// Go语言中函数支持多返回值，函数如果有多个返回值时必须用()将所有返回值包裹起来。
// func calc(x, y int) (int, int) {
// 	sum := x + y
// 	sub := x - y
// 	return sum, sub
// }

// 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过return关键字返回。
// func calc(x, y int) (sum, sub int) {
// 	sum = x + y
// 	sub = x - y
// 	return
// }

// 当我们的一个函数返回值类型为slice时，nil可以看做是一个有效的slice，没必要显示返回一个长度为0的切片。
// func someFunc(x string) []int {
// 	if x == "" {
// 		return nil // 没必要返回[]int{}
// 	}
// 	...
// }

// 函数进阶
// 变量作用域
// 全局变量
// 全局变量是定义在函数外部的变量，它在程序整个运行周期内都有效。 在函数中可以访问到全局变量。

//定义全局变量num
// var num int64 = 10

// func testGlobalVar() {
// 	fmt.Printf("num=%d\n", num) //函数中可以访问全局变量num
// }

// 局部变量
// 局部变量又分为两种： 函数内定义的变量无法在该函数外使用，例如下面的示例代码main函数中无法使用testLocalVar函数中定义的变量x：
// func testLocalVar1() {
// 	//定义一个函数局部变量x,仅在该函数内生效
// 	var x int64 = 100
// 	fmt.Printf("x=%d\n", x)
// }

// 如果局部变量和全局变量重名，优先访问局部变量。
//定义全局变量num
// var num int64 = 10

// func testNum() {
// 	num := 100
// 	fmt.Printf("num=%d\n", num) // 函数中优先使用局部变量
// }

// 接下来我们来看一下语句块定义的变量，通常我们会在if条件判断、for循环、switch语句上使用这种定义变量的方式。
// func testLocalVar2(x, y int) {
// 	fmt.Println(x, y) //函数的参数也是只在本函数中生效
// 	if x > 0 {
// 		z := 100 //变量z只在if语句块生效
// 		fmt.Println(z)
// 	}
// 	//fmt.Println(z)//此处无法使用变量z
// }

// 还有我们之前讲过的for循环语句中定义的变量，也是只在for语句块中生效：
// func testLocalVar3() {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i) //变量i只在当前for语句块中生效
// 	}
// 	//fmt.Println(i) //此处无法使用变量i
// }

// 函数类型与变量
// 定义函数类型
// 我们可以使用type关键字来定义一个函数类型，具体格式如下：

// type calculation func(int, int) int

// 上面语句定义了一个calculation类型，它是一种函数类型，这种函数接收两个int类型的参数并且返回一个int类型的返回值。

// 简单来说，凡是满足这个条件的函数都是calculation类型的函数，例如下面的add和sub是calculation类型。
// func add(x, y int) int {
// 	return x + y
// }

// func sub(x, y int) int {
// 	return x - y
// }

// add和sub都能赋值给calculation类型的变量。

// var c calculation
// c = add

// 高阶函数
// 高阶函数分为函数作为参数和函数作为返回值两部分。
// 函数作为参数
// 函数可以作为参数：

func add(x, y int) int {
	return x + y
}
func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}
