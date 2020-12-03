package main

/*
 * @Author: zut.wxz
 * @Date: 2020-12-03 21:08:03
 * @LastEditors: zut.wxz
 * @LastEditTime: 2020-12-03 21:21:10
 * @Description:
 */

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}
