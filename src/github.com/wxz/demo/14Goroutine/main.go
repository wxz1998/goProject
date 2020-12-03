package main

import "fmt"

/*
 * @Author: zut.wxz
 * @Date: 2020-12-03 21:08:03
 * @LastEditors: zut.wxz
 * @LastEditTime: 2020-12-03 21:34:08
 * @Description:
 */

/*
func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}
*/

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func main() {
	ch := make(chan int)
	go recv(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}
