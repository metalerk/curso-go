package main

import "fmt"

type User interface{
	
	Permisos() int

	Nombre() string

}

type Admin struct{

	nombre string

}

func (self Admin) Permisos() int{
	return 5
}

func (self Admin) Nombre() string{
	return self.nombre
}

func auth(user User) string{
	if user.Permisos() > 4 {
		return user.Nombre() + " tiene permisos de administrador !!"
	}
	return ""
}

func main() {

	admin := Admin{"Luis"}
	fmt.Println(auth(admin))
	
}