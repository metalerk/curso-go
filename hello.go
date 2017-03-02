package main

import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

func main() {
	
	fmt.Println("Hola Mundo !!")

	fmt.Printf("Tu nombre: ")
	
	reader := bufio.NewReader(os.Stdin)
	
	name, err := reader.ReadString('\n')

	if (err != nil){
		fmt.Println("Error: ", err)
	}else{
		fmt.Println("Tu nombre es: ", name)
	}

	//var edad int

	fmt.Printf("Tu edad: ")

	//fmt.Scanf("%d", &edad)
	edad, err2 := reader.ReadString('\n')
	_ = err2

	edad = strings.Split(edad, "\n")[0]

	edad_int,_ := strconv.Atoi(edad)

	if (edad_int < 1){
		fmt.Println("Edad invalida")
	}else if (edad_int > 17) {
		fmt.Println("Eres mayor de edad:",edad_int)
	}else{
		fmt.Println("Eres menor de edad:",edad_int)
	}

}