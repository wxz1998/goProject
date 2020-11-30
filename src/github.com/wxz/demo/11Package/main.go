import "main"

// 包介绍
// 包（package）是多个Go源码的集合，是一种高级的代码复用方案，Go语言为我们提供了很多内置包，如fmt、os、io等。

// 定义包
// 我们还可以根据自己的需要创建自己的包。一个包可以简单理解为一个存放.go文件的文件夹。
// 该文件夹下面的所有go文件都要在代码的第一行添加如下代码，声明该文件归属的包。

// package 包名
// 注意事项：

// 一个文件夹下面直接包含的文件只能归属一个package，同样一个package的文件不能在多个文件夹下。
// 包名可以不和文件夹的名字一样，包名不能包含 - 符号。
// 包名为main的包为应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含main包的源代码则不会得到可执行文件。

// 可见性
// 如果想在一个包中引用另外一个包里的标识符（如变量、常量、类型、函数等）时，该标识符必须是对外可见的（public）。
// 在Go语言中只需要将标识符的首字母大写就可以让标识符对外可见了。

// 举个例子， 我们定义一个包名为pkg2的包，代码如下：

// package pkg2

// import "fmt"

// // 包变量可见性

// var a = 100 // 首字母小写，外部包不可见，只能在当前包内使用

// // 首字母大写外部包可见，可在其他包中使用
// const Mode = 1

// type person struct { // 首字母小写，外部包不可见，只能在当前包内使用
// 	name string
// }

// // 首字母大写，外部包可见，可在其他包中使用
// func Add(x, y int) int {
// 	return x + y
// }

// func age() { // 首字母小写，外部包不可见，只能在当前包内使用
// 	var Age = 18 // 函数局部变量，外部包不可见，只能在当前函数内使用
// 	fmt.Println(Age)
// }

// 单行导入方式定义别名：

// import "fmt"
// import m "github.com/Q1mi/studygo/pkg_test"

// func main() {
// 	fmt.Println(m.Add(100, 200))
// 	fmt.Println(m.Mode)
// }
// 多行导入方式定义别名：

// import (
//     "fmt"
//     m "github.com/Q1mi/studygo/pkg_test"
//  )

// func main() {
// 	fmt.Println(m.Add(100, 200))
// 	fmt.Println(m.Mode)
// }

func main() {
	fmt.Println()
}