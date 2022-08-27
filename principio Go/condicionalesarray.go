package main

import "fmt"

func main (){
	i :=1000


	if i<1000 {
		fmt.Println("el numere es menor que 1000")
	}else{
		fmt.Println("es igual o mayor a 1000")
	}

	foodarray :=[3]string{"pizza","carne","pollo"}/*si sacamos el numero del corchete es un array dinamico*/

//slice es para hacer un hijo del array en este ejemplo hacemos que el hijo solo tenga los 2 primwros ejemplos
foodSlice:= foodarray[0:2] 


for _, food:= range foodarray{/*el _ significa que no vamos a udar el indice para que se ejecute todo el array*/
	switch food{
	case "pizza":
		fmt.Println("vas a comer pizza");
	case "carne":
		fmt.Println("vas a comer carne")
	case "pollo":
		fmt.Println("vas a comer pollo")
	default: 
	fmt.Println("no vas a comer nada")
	}
	}
}

	
