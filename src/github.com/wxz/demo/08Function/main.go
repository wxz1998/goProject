package main

import "fmt"

func main() {
	sayHello()
	ret := intSum(10, 20)
	fmt.Println(ret)
}

func sayHello() {
	fmt.Println("hello")
}

func intSum(x int, y int) int { // 可以写成"x , y int"类型简写
	return x + y
}
