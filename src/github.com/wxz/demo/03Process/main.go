package main

import "fmt"

/* Go语言规定与if匹配的左括号{必须与if和表达式放在同一行，{放在其他位置会触发编译错误。
同理，与else匹配的{也必须与else写在同一行，else也必须与上一个if或else if右边的大括号在同一行。 */
func main() {
	// ifDemo1()
	// ifDemo2()
	// for(循环结构)
	// for 初始语句;条件表达式;结束语句{
	// 循环体语句
	// forDemo()
	switchDemo1()
}

// continue(继续下次循环)
// continue语句可以结束当前循环，开始下一次的循环迭代过程，仅限在for循环内使用。

// 在 continue语句后添加标签时，表示开始标签对应的循环。例如：

// func continueDemo() {
// forloop1:
// 	for i := 0; i < 5; i++ {
// 		// forloop2:
// 		for j := 0; j < 5; j++ {
// 			if i == 2 && j == 2 {
// 				continue forloop1
// 			}
// 			fmt.Printf("%v-%v\n", i, j)
// 		}
// 	}
// }

// break(跳出循环)
// break语句可以结束for、switch和select的代码块。

// break语句还可以在语句后面添加标签，表示退出某个标签对应的代码块，标签要求必须定义在对应的for、switch和 select的代码块上。 举个例子：

// func breakDemo1() {
// BREAKDEMO1:
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			if j == 2 {
// 				break BREAKDEMO1
// 			}
// 			fmt.Printf("%v-%v\n", i, j)
// 		}
// 	}
// 	fmt.Println("...")
// }

// 使用goto语句能简化代码：

// func gotoDemo2() {
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			if j == 2 {
// 				// 设置退出标签
// 				goto breakTag
// 			}
// 			fmt.Printf("%v-%v\n", i, j)
// 		}
// 	}
// 	return
// 	// 标签
// breakTag:
// 	fmt.Println("结束for循环")
// }

// goto(跳转到指定标签)
// goto语句通过标签进行代码间的无条件跳转。goto语句可以在快速跳出循环、避免重复退出上有一定的帮助。Go语言中使用goto语句能简化一些代码的实现过程。 例如双层嵌套的for循环要退出时：

// func gotoDemo1() {
// 	var breakFlag bool
// 	for i := 0; i < 10; i++ {
// 		for j := 0; j < 10; j++ {
// 			if j == 2 {
// 				// 设置退出标签
// 				breakFlag = true
// 				break
// 			}
// 			fmt.Printf("%v-%v\n", i, j)
// 		}
// 		// 外层for循环判断
// 		if breakFlag {
// 			break
// 		}
// 	}
// }

// fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。

// func switchDemo5() {
// 	s := "a"
// 	switch {
// 	case s == "a":
// 		fmt.Println("a")
// 		fallthrough
// 	case s == "b":
// 		fmt.Println("b")
// 	case s == "c":
// 		fmt.Println("c")
// 	default:
// 		fmt.Println("...")
// 	}
// }

// 分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量。例如：

// func switchDemo4() {
// 	age := 30
// 	switch {
// 	case age < 25:
// 		fmt.Println("好好学习吧")
// 	case age > 25 && age < 35:
// 		fmt.Println("好好工作吧")
// 	case age > 60:
// 		fmt.Println("好好享受吧")
// 	default:
// 		fmt.Println("活着真好")
// 	}
// }

// Go语言规定每个switch只能有一个default分支。

// 一个分支可以有多个值，多个case值中间使用英文逗号分隔。

// func testSwitch3() {
// 	switch n := 7; n {
// 	case 1, 3, 5, 7, 9:
// 		fmt.Println("奇数")
// 	case 2, 4, 6, 8:
// 		fmt.Println("偶数")
// 	default:
// 		fmt.Println(n)
// 	}
// }

func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}

// switch case

// for range(键值循环)
// Go语言中可以使用for range遍历数组、切片、字符串、map 及通道（channel）。 通过for range遍历的返回值有以下规律：

// 数组、切片、字符串返回索引和值。
// map返回键和值。
// 通道（channel）只返回通道内的值。

// 无限循环
// for {
//     循环体语句
// }

// for循环的初始语句和结束语句都可以省略，例如：

// func forDemo3() {
// 	i := 0
// 	for i < 10 {
// 		fmt.Println(i)
// 		i++
// 	}
// }

// for循环的初始语句可以被忽略，但是初始语句后的分号必须要写，例如：

// func forDemo2() {
// 	i := 0
// 	for ; i < 10; i++ {
// 		fmt.Println(i)
// 	}
// }

func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

// if条件判断还有一种特殊的写法，可以在 if 表达式之前添加一个执行语句，再根据变量值进行判断
// 优点是变量在if结束后便会被回收,不占用内存 作用域
// func ifDemo2() {
// 	if score := 65; score >= 90 {
// 		fmt.Println("A")
// 	} else if score > 75 {
// 		fmt.Println("B")
// 	} else {
// 		fmt.Println("C")
// 	}
// }

// func ifDemo1() {
// 	score := 65
// 	if score >= 90 {
// 		fmt.Println("A")
// 	} else if score > 75 {
// 		fmt.Println("B")
// 	} else {
// 		fmt.Println("C")
// 	}
// }

// if 表达式1 {
//     分支1
// } else if 表达式2 {
//     分支2
// } else{
//     分支3
// }
