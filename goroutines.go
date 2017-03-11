package main

import (
	"fmt"
	"time"
	"strings"
)

func main() {
	//go mi_nombre_lentooo("Luis")

	go func(){
		letras := strings.Split("1234", "")

	for _,letra := range(letras){
		fmt.Println(letra)
		time.Sleep(1000 * time.Millisecond)
	}
	}()

	fmt.Println("Booooring....")
	var wait string
	fmt.Scanln(&wait)
}

func mi_nombre_lentooo(name string){
	letras := strings.Split(name, "")

	for _,letra := range(letras){
		fmt.Println(letra)
		time.Sleep(1000 * time.Millisecond)
	}
}