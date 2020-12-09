package main

import "fmt"

// go语言函数外的语句必须以关键字开头
var name = "wxz"

// 常量声明
const pi = 3.1415

// 多个初始化
var name1, age1 = "bbb", 22

// 变量的声明方式
func main() {
	// 赋值操作只能在函数体内
	age := 18
	fmt.Println(name, age, pi)
	fmt.Println(name1, age1)
}
