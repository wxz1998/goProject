package main

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
type Mover interface {
	move()
}

type dog struct{}
