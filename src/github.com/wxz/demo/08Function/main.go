package main

import "fmt"

func main() {
	sayHello()
	ret := intSum(10, 20)
	fmt.Println(ret)

	ret1 := intSum2()
	ret2 := intSum2(10)
	ret3 := intSum2(10, 20)
	ret4 := intSum2(10, 20, 30)
	fmt.Println(ret1, ret2, ret3, ret4) //0 10 30 60

	ret5 := intSum3(100)
	ret6 := intSum3(100, 10)
	ret7 := intSum3(100, 10, 20)
	ret8 := intSum3(100, 10, 20, 30)
	fmt.Println(ret5, ret6, ret7, ret8) //100 110 130 160
}

func sayHello() {
	fmt.Println("hello")
}

func intSum(x int, y int) int { // 可以写成"x , y int"类型简写
	return x + y
}

// 可变参数是指函数的参数数量不固定。Go语言中的可变参数通过在参数名后加...来标识。
// 注意：可变参数通常要作为函数的最后一个参数。
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

// 固定参数搭配可变参数使用时，可变参数要放在固定参数的后面
func intSum3(x int, y ...int) int {
	fmt.Println(x, y)
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}

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
