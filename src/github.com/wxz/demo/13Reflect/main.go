package main

import (
	"fmt"
	"reflect"
)

/*
 * @Author: zut.wxz
 * @Date: 2020-12-02 15:12:58
 * @LastEditors: zut.wxz
 * @LastEditTime: 2020-12-02 16:50:12
 * @Description:
 */

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type:%v\n", v)
}
func main() {
	var a float32 = 3.14
	reflectType(a) // type:float32
	var b int64 = 100
	reflectType(b) // type:int64
}
