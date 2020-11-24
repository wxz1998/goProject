package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}
	s := a[1:3] // s := a[low:high]
	fmt.Printf("a:%v s:%v len(s):%v cap(s):%v\n", a, s, len(s), cap(s))

	// a := [...]float64{67.7, 89.8, 21, 78}
	// sum := float64(0)
	// for i, v := range a { //range returns both the index and value
	// 	fmt.Printf("%d the element of a is %.2f\n", i, v)
	// 	sum += v
	// }
	// fmt.Println("\nsum of all elements of a", sum)
}
