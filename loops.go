package main

import "fmt"

func main() {
	count := 10
	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	fmt.Println("----------")

	i := 10

	for i != 0 {
		fmt.Println(i)
		i --
	}
}