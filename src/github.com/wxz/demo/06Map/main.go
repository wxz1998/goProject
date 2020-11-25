package main

import "fmt"

func main() {
	// a := [5]int{1, 2, 3, 4, 5}
	// s := a[1:3] // s := a[low:high]
	// fmt.Printf("a:%v s:%v len(s):%v cap(s):%v\n", a, s, len(s), cap(s))

	// 使用range遍历数组
	// a := [...]float64{67.7, 89.8, 21, 78}
	// sum := float64(0)
	// for _, v := range a { //range returns both the index and value
	// 	fmt.Printf("the element of a is %.2f\n", v)
	// 	sum += v
	// }
	// fmt.Println("\nsum of all elements of a", sum)

	// 多维数组
	// 声明多维数组的方法一
	// a := [3][2]string{
	// 	{"lion", "tiger"},
	// 	{"cat", "dog"},
	// 	{"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
	// }
	// printarray(a)
	// 声明多维数组的方法二
	// var b [3][2]string
	// b[0][0] = "apple"
	// b[0][1] = "samsung"
	// b[1][0] = "microsoft"
	// b[1][1] = "google"
	// b[2][0] = "AT&T"
	// b[2][1] = "T-Mobile"
	// fmt.Printf("\n")
	// printarray(b)

	// 切片
	// 创建切片的方法一
	// // 先创建数组
	// a := [5]int{76, 77, 78, 79, 80}
	// // 从数组a中拿出a[1]到a[3]的数据放入切片b 左包含 右不包含 [)
	// var b []int = a[1:4] //creates a slice from a[1] to a[3]
	// fmt.Println(b)

	// // 创建切片的方法二
	// // 创建数组并返回给切片c 但怎么看都像数组...
	// c := []int{6, 7, 8} //creates and array and returns a slice reference
	// fmt.Println(c)
	// // 对照创建数组 有什么不同?
	// // a := [3]int{5, 78, 8}
	// // a := [...]int{12, 78, 50}
	// // a := [...]string{"USA", "China", "India", "Germany", "France"}

	// // 切片反作用于数组
	// // 创建数组
	// darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	// // 从数组中获得切片
	// dslice := darr[2:5]
	// fmt.Println("array before", darr)
	// for i := range dslice {
	// 	dslice[i]++
	// 	// 从数组中获得的切片能反影响到数组本身 切片 下层数据的上层表现
	// }
	// fmt.Println("array after", darr)

	// 当若干个切片共享同一个底层数组时，对每一个切片的修改都会反映在底层数组中。
	// numa := [3]int{78, 79, 80}
	// 默认全切
	// nums1 := numa[:] //creates a slice which contains all elements of the array
	// nums2 := numa[:]
	// fmt.Println("array before change 1", numa)
	// nums1[0] = 100
	// fmt.Println("array after modification to slice nums1", numa)
	// nums2[1] = 101
	// fmt.Println("array after modification to slice nums2", numa)
	// 从输出结果可以看出，当多个切片共享同一个数组时，对每一个切片的修改都将会反映到这个数组中。

	// // 切片的长度和容量
	// fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	// fruitslice := fruitarray[1:3]
	// fmt.Println(fruitslice)
	// fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
	// // fruitarray 的长度是 7。fruiteslice 是从 fruitarray 的索引 1 开始的。
	// // 因此 fruiteslice 的容量是从 fruitarray 的第 1 个元素开始算起的数组中的元素个数，这个值是 6。因此 fruitslice 的容量是 6。
	// // 切片的长度可以动态的改变（最大为其容量）。任何超出最大容量的操作都会发生运行时错误。

	// // 修改 fruitslice 的长度为它的容量。
	// fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
	// fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))

	// 	用 make 创建切片
	// 内置函数 func make([]T, len, cap) []T 可以用来创建切片，该函数接受长度和容量作为参数，返回切片。
	// 容量是可选的，默认与长度相同。使用 make 函数将会创建一个数组并返回它的切片。
	i := make([]int, 5, 5)
	fmt.Println(i)
	// 用 make 创建的切片的元素值默认为 0 值。上面的程序输出为：[0 0 0 0 0]。
}

// 使用2个range,嵌套的方式打印多维数组
// func printarray(a [3][2]string) {
// 	for _, v1 := range a {
// 		for _, v2 := range v1 {
// 			fmt.Printf("%s ", v2)
// 		}
// 		fmt.Printf("\n")
// 	}
// }
