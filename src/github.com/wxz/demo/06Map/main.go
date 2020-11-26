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
	// i := make([]int, 5, 5)
	// fmt.Println(i)
	// 用 make 创建的切片的元素值默认为 0 值。上面的程序输出为：[0 0 0 0 0]。

	// 追加元素到切片
	// 数组是固定长度的，它们的长度不能动态增加。而切片是动态的，可以使用内置函数 append 添加元素到切片。
	// append 的函数原型为：append(s []T, x ...T) []T。
	// x …T 表示 append 函数可以接受的参数个数是可变的。这种函数叫做变参函数。

	// 问题：如果切片是建立在数组之上的，而数组本身不能改变长度，那么切片是如何动态改变长度的呢？
	// 实际发生的情况是，当新元素通过调用 append 函数追加到切片末尾时，如果超出了容量，append 内部会创建一个新的数组。
	// 并将原有数组的元素被拷贝给这个新的数组，最后返回建立在这个新数组上的切片。
	// 这个新切片的容量是旧切片的二倍（译者注：当超出切片的容量时，append 将会在其内部创建新的数组，该数组的大小是原切片容量的 2 倍。
	// 最后 append 返回这个数组的全切片，即从 0 到 length - 1 的切片）。
	// cars := []string{"Ferrari", "Honda", "Ford"}
	// fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	// cars = append(cars, "Toyota")
	// fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	// 切片的 0 值为 nil。一个 nil 切片的长度和容量都为 0。可以利用 append 函数给一个 nil 切片追加值。

	// var names []string //zero value of a slice is nil
	// if names == nil {
	// 	fmt.Println("slice is nil going to append")
	// 	names = append(names, "John", "Sebastian", "Vinay")
	// 	fmt.Println("names contents:", names)
	// }

	// 可以使用 ... 操作符将一个切片追加到另一个切片末尾：
	// veggies := []string{"potatoes", "tomatoes", "brinjal"}
	// fruits := []string{"oranges", "apples"}
	// food := append(veggies, fruits...)
	// fmt.Println("food:", food)

	// 	切片作为函数参数
	// 可以认为切片在内部表示为如下的结构体：
	// type slice struct {
	//     Length        int
	//     Capacity      int
	//     ZerothElement *byte
	// }
	// 可以看到切片包含长度、容量、以及一个指向首元素的指针。
	// 当将一个切片作为参数传递给一个函数时，虽然是值传递，但是指针始终指向同一个数组。
	// 因此将切片作为参数传给函数时，函数对该切片的修改在函数外部也可以看到。

	// nos := []int{8, 7, 6}
	// fmt.Println("slice before function call", nos)
	// subtactOne(nos)                               //function modifies the slice
	// fmt.Println("slice after function call", nos) //modifications are visible outside

	// 	多维切片
	// 同数组一样，切片也可以有多个维度。
	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
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

// func subtactOne(numbers []int) {
// 	for i := range numbers {
// 		numbers[i] -= 2
// 	}

// }
