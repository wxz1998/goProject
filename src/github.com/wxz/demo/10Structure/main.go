package main

import "fmt"

func main() {
	fmt.Println()
}

// 自定义类型
// 在Go语言中有一些基本的数据类型，如string、整型、浮点型、布尔等数据类型， Go语言中可以使用type关键字来定义自定义类型。

// 自定义类型是定义了一个全新的类型。我们可以基于内置的基本类型定义，也可以通过struct定义。例如：

//将MyInt定义为int类型
// type MyInt int

// 通过type关键字的定义，MyInt就是一种新的类型，它具有int的特性。

// 类型别名
// 类型别名是Go1.9版本添加的新功能。

// 类型别名规定：TypeAlias只是Type的别名，本质上TypeAlias与Type是同一个类型。就像一个孩子小时候有小名、乳名，上学后用学名，英语老师又会给他起英文名，但这些名字都指的是他本人。

// type TypeAlias = Type
// 我们之前见过的rune和byte就是类型别名，他们的定义如下：

// type byte = uint8
// type rune = int32
