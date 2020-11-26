package main

import "fmt"

func main() {
	// Go语言中 map的定义语法如下：
	// map[KeyType]ValueType
	// 	KeyType:表示键的类型。
	// ValueType:表示键对应的值的类型。

	// 	map类型的变量默认初始值为nil，需要使用make()函数来分配内存。语法为：
	// make(map[KeyType]ValueType, [cap])
	// 其中cap表示map的容量，该参数虽然不是必须的，但是我们应该在初始化map的时候就为其指定一个合适的容量。

	//map基本使用
	// map中的数据都是成对出现的，map的基本使用示例代码如下：
	// map内参数成对出现 下面 string 指姓名 int 指分数
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)
}
