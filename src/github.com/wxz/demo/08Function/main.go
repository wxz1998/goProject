package main

import "fmt"

// func main() {
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

// }

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

// ##############################################################
// 高阶函数
// 高阶函数分为函数作为参数和函数作为返回值两部分。
// 函数作为参数
// 函数可以作为参数：

// func add(x, y int) int {
// 	return x + y
// }
// func calc(x, y int, op func(int, int) int) int {
// 	return op(x, y)
// }

// func main() {
// 	ret2 := calc(10, 20, add)
// 	fmt.Println(ret2) //30
// }

// 函数作为返回值
// 函数也可以作为返回值：
// func do(s string) (func(int, int) int, error) {
// 	switch s {
// 	case "+":
// 		return add, nil
// 	case "-":
// 		return sub, nil
// 	default:
// 		err := errors.New("无法识别的操作符")
// 		return nil, err
// 	}
// }

// 匿名函数和闭包
// 匿名函数
// 函数当然还可以作为返回值，但是在Go语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。匿名函数就是没有函数名的函数，匿名函数的定义格式如下：

// func(参数)(返回值){
//     函数体
// }

// 匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数:

// func main() {
// 	// 将匿名函数保存到变量
// 	add := func(x, y int) {
// 		fmt.Println(x + y)
// 	}
// 	add(10, 20) // 通过变量调用匿名函数

// 	//自执行函数：匿名函数定义完加()直接执行
// 	func(x, y int) {
// 		fmt.Println(x + y)
// 	}(10, 20)
// }

// 匿名函数多用于实现回调函数和闭包。

// 闭包
// 闭包指的是一个函数和与其相关的引用环境组合而成的实体。简单来说，闭包=函数+引用环境。 首先我们来看一个例子：

// func adder() func(int) int {
// 	var x int
// 	return func(y int) int {
// 		x += y
// 		return x
// 	}
// }
// func main() {
// 	var f = adder()
// 	fmt.Println(f(10)) //10
// 	fmt.Println(f(20)) //30
// 	fmt.Println(f(30)) //60

// 	f1 := adder()
// 	fmt.Println(f1(40)) //40
// 	fmt.Println(f1(50)) //90
// }

// 变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。
// 在f的生命周期内，变量x也一直有效。 闭包进阶示例1：
func adder2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
	var f = adder2(10)
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70

	f1 := adder2(20)
	fmt.Println(f1(40)) //60
	fmt.Println(f1(50)) //110
}

// 闭包进阶示例2：

// func makeSuffixFunc(suffix string) func(string) string {
// 	return func(name string) string {
// 		if !strings.HasSuffix(name, suffix) {
// 			return name + suffix
// 		}
// 		return name
// 	}
// }

// func main() {
// 	jpgFunc := makeSuffixFunc(".jpg")
// 	txtFunc := makeSuffixFunc(".txt")
// 	fmt.Println(jpgFunc("test")) //test.jpg
// 	fmt.Println(txtFunc("test")) //test.txt
// }

// 闭包进阶示例3：

// func calc(base int) (func(int) int, func(int) int) {
// 	add := func(i int) int {
// 		base += i
// 		return base
// 	}

// 	sub := func(i int) int {
// 		base -= i
// 		return base
// 	}
// 	return add, sub
// }

// func main() {
// 	f1, f2 := calc(10)
// 	fmt.Println(f1(1), f2(2)) //11 9
// 	fmt.Println(f1(3), f2(4)) //12 8
// 	fmt.Println(f1(5), f2(6)) //13 7
// }
// 闭包其实并不复杂，只要牢记闭包=函数+引用环境。

// defer语句
// Go语言中的defer语句会将其后面跟随的语句进行延迟处理。
// 在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，
// 也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。

// func main() {
// 	fmt.Println("start")
// 	defer fmt.Println(1)
// 	defer fmt.Println(2)
// 	defer fmt.Println(3)
// 	fmt.Println("end")
// }

// 由于defer语句延迟调用的特性，所以defer语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。

// defer经典案例
// 阅读下面的代码，写出最后的打印结果。
