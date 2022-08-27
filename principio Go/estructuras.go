package main

import "fmt"

type Person struct{
	name string
	email string
	age int
}

func main(){
	person := Person{
		"Tomas",
	"corretomi",
		 20,
	}
	person.name = "agus"

	fmt.Println("El nombre de la persona es: ",person.name, " su correo es : ",person.email, " la edad es :", person.age)
}