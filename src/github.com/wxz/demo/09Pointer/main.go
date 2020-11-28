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
// func main() {
// 	//指针取值
// 	a := 10
// 	b := &a                          // 取变量a的地址，将指针保存到b中
// 	fmt.Printf("type of b:%T\n", b)  // type of b:*int
// 	c := *b                          // 指针取值（根据指针去内存取值）
// 	fmt.Printf("type of c:%T\n", c)  // type of c:int
// 	fmt.Printf("value of c:%v\n", c) // value of c:10
// }

// 总结： 取地址操作符&和取值操作符*是一对互补操作符，&取出地址，*根据地址取出地址指向的值。

// 变量、指针地址、指针变量、取地址、取值的相互关系和特性如下：

// 对变量进行取地址（&）操作，可以获得这个变量的指针变量。
// 指针变量的值是指针地址。
// 对指针变量进行取值（*）操作，可以获得指针变量指向的原变量的值。

// 指针传值示例：

// func modify1(x int) {
// 	x = 100
// }

// func modify2(x *int) {
// 	*x = 100
// }

// func main() {
// 	a := 10
// 	modify1(a)
// 	fmt.Println(a) // 10
// 	modify2(&a)
// 	fmt.Println(a) // 100
// }

// new和make

func main() {
	var a *int
	*a = 100
	fmt.Println(*a)

	var b map[string]int
	b["沙河娜扎"] = 100
	fmt.Println(b)
}

// 执行上面的代码会引发panic，为什么呢？
// 在Go语言中对于引用类型的变量，我们在使用的时候不仅要声明它，还要为它分配内存空间，否则我们的值就没办法存储。
// 而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。
// 要分配内存，就引出来今天的new和make。
// Go语言中new和make是内建的两个函数，主要用来分配内存。

// new
// new是一个内置的函数，它的函数签名如下：

// func new(Type) *Type

// 其中，

// Type表示类型，new函数只接受一个参数，这个参数是一个类型
// *Type表示类型指针，new函数返回一个指向该类型内存地址的指针。
