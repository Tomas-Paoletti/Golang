package main

import "fmt"

func printTextFunction(name string){
	fmt.Println("Hello wordl func",name, getAge())
}
func getAge()int{
	var age int
	fmt.Println("cual es tu edad")
	fmt.Scan(&age)
	return age
}

func main(){
	var name string
	fmt.Println("Escribi tu nombre")
	fmt.Scan(&name)

	printTextFunction(name, )

	
}