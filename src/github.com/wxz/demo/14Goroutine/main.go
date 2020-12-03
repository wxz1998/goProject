package main

/*
 * @Author: zut.wxz
 * @Date: 2020-12-03 21:08:03
 * @LastEditors: zut.wxz
 * @LastEditTime: 2020-12-03 21:19:14
 * @Description:
 */

import "fmt"

func hello() {
	fmt.Println("hello goroutine")
}

func main() {
	go hello()
	fmt.Println("main goroutine done")
}
