package main

import "fmt"
import "strconv"

func main() {

	b_str := "true"

	b, err := strconv.ParseBool(b_str)

	if err == nil {

		fmt.Println(b)

	}else{
		fmt.Println(err)
	}
	
}