package main

import (
	"fmt"
	"math"
)

func main() {
	// // 十进制
	// var a int = 10
	// fmt.Printf("%d \n", a) // 10
	// fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// // 八进制  以0开头
	// var b int = 077
	// fmt.Printf("%o \n", b) // 77

	// // 十六进制  以0x开头
	// var c int = 0xff
	// fmt.Printf("%x \n", c) // ff
	// fmt.Printf("%X \n", c) // FF

	// fmt.Printf("%f\n", math.Pi)
	// fmt.Printf("%.2f\n", math.Pi)

	// //复数
	// var c1 complex64
	// c1 = 1 + 2i
	// var c2 complex128
	// c2 = 2 + 3i
	// fmt.Println(c1)
	// fmt.Println(c2)

	// // 打印一个路径
	// fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

	// 	// 用反引号定义多行字符串
	// 	s1 := `第一行
	// 第二行
	// 第三行
	// `
	// 	fmt.Println(s1)

	// // 遍历字符串
	// traversalString()

	// // 修改字符串
	// changeString()

	// 类型转换
	sqrtDemo()
}

// // 遍历字符串
// func traversalString() {
// 	s := "hello沙河"
// 	for i := 0; i < len(s); i++ { //byte
// 		fmt.Printf("%v(%c) ", s[i], s[i])
// 	}
// 	fmt.Println()
// 	for _, r := range s { //rune
// 		fmt.Printf("%v(%c) ", r, r)
// 	}
// 	fmt.Println()
// }

// // 修改字符串
// func changeString() {
// 	s1 := "big"
// 	// 强制类型转换
// 	byteS1 := []byte(s1)
// 	byteS1[0] = 'p'
// 	fmt.Println(string(byteS1))

// 	s2 := "白萝卜"
// 	runeS2 := []rune(s2)
// 	runeS2[0] = '红'
// 	fmt.Println(string(runeS2))
// }

// 类型转换
func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}
