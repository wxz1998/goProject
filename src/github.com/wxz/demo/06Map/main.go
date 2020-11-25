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

	// 创建数组
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	// 从数组中获得切片
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
		// 从数组中获得的切片能反影响到数组本身 切片 下层数据的上层表现
	}
	fmt.Println("array after", darr)
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
