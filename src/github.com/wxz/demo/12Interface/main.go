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
// 	var peo People = Student{}
// 	think := "bitch"
// 	fmt.Println(peo.Speak(think))
// }
// 	var peo People = Student{}
// 	think := "bitch"
// 	fmt.Println(peo.Speak(think))
// }
// 	var peo People = Student{}
// 	think := "bitch"
// 	fmt.Println(peo.Speak(think))
// }
// 	var peo People = Student{}
// 	think := "bitch"
// 	fmt.Println(peo.Speak(think))
// }
// 	var peo People = Student{}
// 	think := "bitch"
// 	fmt.Println(peo.Speak(think))
// }
// 	var peo People = Student{}
// 	think := "bitch"
// 	fmt.Println(peo.Speak(think))
// }
// 	var peo People = Student{}
// 	think := "bitch"
// 	fmt.Println(peo.Speak(think))
// }
