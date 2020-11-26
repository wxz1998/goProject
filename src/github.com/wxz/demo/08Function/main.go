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
}

func sayHello() {
	fmt.Println("hello")
}

func intSum(x int, y int) int { // 可以写成"x , y int"类型简写
	return x + y
}

func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
