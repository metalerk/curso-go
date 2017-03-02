package main

import "fmt"

func main() {
	
	var x,y *int

	entero := 5

	x = &entero
	y = &entero

	*x = 6

	fmt.Println(*x)
	fmt.Println(*y)

	slice := []int{1,2}

	fmt.Printf("%p\n", &slice)

}