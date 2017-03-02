package main

import "fmt"

func main() {
	
	//slice := make([]int, 3, 5)
	slice := []int{1,2,3,4}
	slice2 := make([]int, len(slice), cap(slice) * 2)

	fmt.Println(slice2)

	
	/*slice = append(slice, 1)

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	*/

	cp := copy(slice2, slice)

	fmt.Println("Copy", cp)

	fmt.Println(slice)

	fmt.Println(slice2)

}