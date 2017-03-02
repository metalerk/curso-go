package main

import "fmt"

func main() {
	
	//var arreglo [10]int
	arreglo := [10]int{1,4}

	fmt.Println(arreglo)

	for i := 0; i < len(arreglo); i++ {
		arreglo[i] = i + 1
	}
	fmt.Println(arreglo)

	var matrix [3][2]int
	matrix[0][1] = 1

	fmt.Println(matrix)

}