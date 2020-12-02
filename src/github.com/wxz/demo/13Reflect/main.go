package main

import (
	"fmt"
	"reflect"
)

/*
 * @Author: zut.wxz
 * @Date: 2020-12-02 15:12:58
 * @LastEditors: zut.wxz
 * @LastEditTime: 2020-12-02 16:49:03
 * @Description:
 */

// 在Go语言中，使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），程序通过类型对象可以访问任意值的类型信息。
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
