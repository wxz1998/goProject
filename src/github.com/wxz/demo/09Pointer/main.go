package main

import "fmt"

// 指针地址和指针类型
// Go语言中的指针不能进行偏移和运算，因此Go语言中的指针操作非常简单，我们只需要记住两个符号：&（取地址）和*（根据地址取值）
// func main() {
// 	a := 10
// 	b := &a                            // 把a的内存地址返回给b
// 	fmt.Printf("a:%d ptr:%p\n", a, &a) // a的值和a的内存地址 a:10 ptr:0xc0000160a0
// 	fmt.Printf("b:%p type:%T\n", b, b) // b的值和类型 b:0xc0000160a0 type:*int
// 	fmt.Println(&b)                    // 变量b的内存地址 0xc000006028
// }

// 指针取值
// 在对普通变量使用&操作符取地址后会获得这个变量的指针，然后可以对指针使用*操作，也就是指针取值
func main() {
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}
