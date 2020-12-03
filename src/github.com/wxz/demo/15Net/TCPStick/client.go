package main

import (
	"fmt"
	"net"
)

/*
 * @Author: zut.wxz
 * @Date: 2020-12-03 23:03:11
 * @LastEditors: zut.wxz
 * @LastEditTime: 2020-12-03 23:03:36
 * @Description:
 */

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		conn.Write([]byte(msg))
	}
}
