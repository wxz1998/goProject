package main

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
