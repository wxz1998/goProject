package main

import "fmt"

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

// 类型定义和类型别名的区别
// 类型别名与类型定义表面上看只有一个等号的差异，我们通过下面的这段代码来理解它们之间的区别。

//类型定义

// type NewInt int

//类型别名

// type MyInt = int

// func main() {
// 	var a NewInt
// 	var b MyInt

// 	fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
// 	fmt.Printf("type of b:%T\n", b) //type of b:int
// }

// 结果显示a的类型是main.NewInt，表示main包下定义的NewInt类型。b的类型是int。MyInt类型只会在代码中存在，编译完成时并不会有MyInt类型。

// 结构体的定义
// 使用type和struct关键字来定义结构体，具体代码格式如下：

// type 类型名 struct {
//     字段名 字段类型
//     字段名 字段类型
//     …
// }
// 其中：

// 类型名：标识自定义结构体的名称，在同一个包内不能重复。
// 字段名：表示结构体字段名。结构体中的字段名必须唯一。
// 字段类型：表示结构体字段的具体类型。

// 我们定义一个Person（人）结构体

// type person struct {
// 	name string
// 	city string
// 	age  int8
// }

// 同样类型的字段也可以写在一行，

// type person1 struct {
// 	name, city string
// 	age        int8
// }

// 这样我们就拥有了一个person的自定义类型，它有name、city、age三个字段，分别表示姓名、城市和年龄。
// 这样我们使用这个person结构体就能够很方便的在程序中表示和存储人信息了。

// 语言内置的基础数据类型是用来描述一个值的，而结构体是用来描述一组值的。
// 比如一个人有名字、年龄和居住城市等，本质上是一种聚合型的数据类型

// 基本实例化
// type person struct {
// 	name string
// 	city string
// 	age  int8
// }

// func main() {
// 	var p1 person
// 	p1.name = "沙河娜扎"
// 	p1.city = "北京"
// 	p1.age = 18
// 	fmt.Printf("p1=%v\n", p1)  //p1={沙河娜扎 北京 18}
// 	fmt.Printf("p1=%#v\n", p1) //p1=main.person{name:"沙河娜扎", city:"北京", age:18}
// }

// 我们通过.来访问结构体的字段（成员变量）,例如p1.name和p1.age等。

// 匿名结构体
// 在定义一些临时数据结构等场景下还可以使用匿名结构体。
func main() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%#v\n", user)
}

// 创建指针类型结构体
// 我们还可以通过使用new关键字对结构体进行实例化，得到的是结构体的地址。
