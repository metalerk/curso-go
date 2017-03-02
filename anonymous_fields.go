package main

import "fmt"

type Human struct{
	name string
}

type Dummy struct{
	name string
}

type Tutor struct{
	Human
	Dummy
}

func (self *Human) hablar() string{
	return "Hola Humano !!"
}

func (self *Tutor) hablar() string{
	return "Hola Tutor !!"
}

func main() {
	
	tutor := Tutor{Human{"Luis"}, Dummy{"Esteban"}}
	
	//fmt.Println(tutor.name)
	fmt.Println(tutor.Human.name)
	fmt.Println(tutor.Dummy.name)
	fmt.Println(tutor.Human.hablar())
	fmt.Println(tutor.hablar())

}