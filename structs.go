package main

import "fmt"

type User struct{

	edad int

	nombre string

	apellido string

}

func (self *User) nombre_completo () string{
	
	return self.nombre + " " + self.apellido

}

func (self *User) set_name(name string){
	self.nombre = name
}

func main() {
	
	//var luis User

	/*
	luis := User{24, "Luis Esteban", "Rodriguez"}

	luis.nombre = "Luis"
	*/

	luis := new(User)

	(*luis).edad = 24

	(*luis).nombre = "Luis Esteban"

	(*luis).apellido = "Rodriguez"

	(*luis).set_name("Jose")

	fmt.Println(luis.nombre_completo())
	
	//fmt.Printf("%p\n%p\n", &luis, &luis.nombre)

}