package main

import "fmt"

// go语言函数外的语句必须以关键字开头
var name = "wxz"

// 常量声明 函数体外声明的变量为全局变量
const pi = 3.1415

// 多个初始化
var name1, age1 = "bbb", 22

// 变量的声明方式
func main() {
	// 局部变量
	age := 18
	fmt.Println(name, age, pi)
	fmt.Println(name1, age1)
	// 哑元变量 也叫匿名变量
	x, _ := test1()
	_, y := test1()
	fmt.Println(x, y)
}

func test1() (int, string) {
	return 13, "ccc"
}
