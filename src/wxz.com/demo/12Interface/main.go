package main

import "fmt"

/*
 * @Author: zut.wxz
 * @Date: 2020-11-30 18:11:19
 * @LastEditors: zut.wxz
 * @LastEditTime: 2020-12-02 15:11:22
 * @Description:
 */

// func main() {
// 	var x Sayer // 声明一个Sayer类型的变量x
// 	a := cat{}  // 实例化一个cat
// 	b := dog{}  // 实例化一个dog
// 	x = a       // 可以把cat实例直接赋值给x
// 	x.say()     // 喵喵喵
// 	x = b       // 可以把dog实例直接赋值给x
// 	x.say()     // 汪汪汪}
// }

// // Sayer 接口
// type Sayer interface {
// 	say()
// }
// type dog struct{}

// type cat struct{}

// // dog实现了Sayer接口
// func (d dog) say() {
// 	fmt.Println("汪汪汪")
// }

// // cat实现了Sayer接口
// func (c cat) say() {
// 	fmt.Println("喵喵喵")
// }

// Tips： 观察下面的代码，体味此处_的妙用

// // 摘自gin框架routergroup.go
// type IRouter interface{ ... }

// type RouterGroup struct { ... }

// var _ IRouter = &RouterGroup{}  // 确保RouterGroup实现了接口IRouter

// 值接收者和指针接收者实现接口的区别
// 使用值接收者实现接口和使用指针接收者实现接口有什么区别呢？接下来我们通过一个例子看一下其中的区别。

// 我们有一个Mover接口和一个dog结构体。

// type Mover interface {
// 	move()
// }

// type dog struct{}

// // 值接收者实现接口
// func (d dog) move() {
// 	fmt.Println("狗会动")
// }

// // 此时实现接口的是dog类型：
// func main() {
// 	var x Mover
// 	var wangcai = dog{} // 旺财是dog类型
// 	x = wangcai         // x可以接收dog类型
// 	var fugui = &dog{}  // 富贵是*dog类型
// 	x = fugui           // x可以接收*dog类型
// 	x.move()
// }

// 从上面的代码中我们可以发现，使用值接收者实现接口之后，不管是dog结构体还是结构体指针*dog类型的变量都可以赋值给该接口变量。
// 因为Go语言中有对指针类型变量求值的语法糖，dog指针fugui内部会自动求值*fugui。

// 指针接收者实现接口
// 同样的代码我们再来测试一下使用指针接收者有什么区别：

// func (d *dog) move() {
// 	fmt.Println("狗会动")
// }
// func main() {
// 	var x Mover
// 	var wangcai = dog{} // 旺财是dog类型
// 	x = wangcai         // x不可以接收dog类型
// 	var fugui = &dog{}  // 富贵是*dog类型
// 	x = fugui           // x可以接收*dog类型
// }

// 此时实现Mover接口的是*dog类型，所以不能给x传入dog类型的wangcai，此时x只能存储*dog类型的值。

// 面试题
// 注意：这是一道你需要回答"能"或者"不能"的题！

// 首先请观察下面的这段代码，然后请回答这段代码能不能通过编译？
// type People interface {
// 	Speak(string) string
// }

// type Student struct{}

// func (stu *Student) Speak(think string) (talk string) {
// 	if think == "sb" {
// 		talk = "你是个大帅比"
// 	} else {
// 		talk = "您好"
// 	}
// 	return
// }

// func main() {
// 	var peo People = Student{}
// 	think := "bitch"
// 	fmt.Println(peo.Speak(think))
// }

// 类型与接口的关系
// 一个类型实现多个接口
// 一个类型可以同时实现多个接口，而接口间彼此独立，不知道对方的实现。 例如，狗可以叫，也可以动。我们就分别定义Sayer接口和Mover接口，如下： Mover接口。

// // Sayer 接口
// type Sayer interface {
// 	say()
// }

// // Mover 接口
// type Mover interface {
// 	move()
// }

// // dog既可以实现Sayer接口，也可以实现Mover接口。
// type dog struct {
// 	name string
// }

// // 实现Sayer接口
// func (d dog) say() {
// 	fmt.Printf("%s会叫汪汪汪\n", d.name)
// }

// // 实现Mover接口
// func (d dog) move() {
// 	fmt.Printf("%s会动\n", d.name)
// }

// func main() {
// 	var x Sayer
// 	var y Mover

// 	var a = dog{name: "旺财"}
// 	x = a
// 	y = a
// 	x.say()
// 	y.move()
// }

// 多个类型实现同一接口
// Go语言中不同的类型还可以实现同一接口 首先我们定义一个Mover接口，它要求必须由一个move方法。

// // Mover 接口
// type Mover interface {
// 	move()
// }

// // 例如狗可以动，汽车也可以动，可以使用如下代码实现这个关系：

// type dog struct {
// 	name string
// }

// type car struct {
// 	brand string
// }

// // dog类型实现Mover接口
// func (d dog) move() {
// 	fmt.Printf("%s会跑\n", d.name)
// }

// // car类型实现Mover接口
// func (c car) move() {
// 	fmt.Printf("%s速度70迈\n", c.brand)
// }

// // 这个时候我们在代码中就可以把狗和汽车当成一个会动的物体来处理了，不再需要关注它们具体是什么，只需要调用它们的move方法就可以了。

// func main() {
// 	var x Mover
// 	var a = dog{name: "旺财"}
// 	var b = car{brand: "保时捷"}
// 	x = a
// 	x.move()
// 	x = b
// 	x.move()
// }

// 接口嵌套
// 接口与接口间可以通过嵌套创造出新的接口。

// // Sayer 接口
// type Sayer interface {
// 	say()
// }

// // Mover 接口
// type Mover interface {
// 	move()
// }

// // 接口嵌套
// type animal interface {
// 	Sayer
// 	Mover
// }

// // 嵌套得到的接口的使用与普通接口一样，这里我们让cat实现animal接口：

// type cat struct {
// 	name string
// }

// func (c cat) say() {
// 	fmt.Println("喵喵喵")
// }

// func (c cat) move() {
// 	fmt.Println("猫会动")
// }

// func main() {
// 	var x animal
// 	x = cat{name: "花花"}
// 	x.move()
// 	x.say()
// }

// 空接口
// 空接口的定义
// 空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。

// 空接口类型的变量可以存储任意类型的变量。

func main() {
	// 定义一个空接口x
	var x interface{}
	s := "Hello 沙河"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)
}

// 空接口的应用
// 空接口作为函数的参数
// 使用空接口实现可以接收任意类型的函数参数。

// 空接口作为函数参数
// func show(a interface{}) {
// 	fmt.Printf("type:%T value:%v\n", a, a)
// }

// 空接口作为map的值
// 使用空接口实现可以保存任意值的字典。

// 空接口作为map值
// var studentInfo = make(map[string]interface{})
// studentInfo["name"] = "沙河娜扎"
// studentInfo["age"] = 18
// studentInfo["married"] = false
// fmt.Println(studentInfo)
