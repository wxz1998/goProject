package main

import "fmt"

// 指针地址和指针类型
func main() {
	a := 10
	b := &a                            // 把a的内存地址返回给b
	fmt.Printf("a:%d ptr:%p\n", a, &a) // a的值和a的内存地址 a:10 ptr:0xc0000160a0
	fmt.Printf("b:%p type:%T\n", b, b) // b的值和类型 b:0xc0000160a0 type:*int
	fmt.Println(&b)                    // 变量b的内存地址 0xc000006028
}
