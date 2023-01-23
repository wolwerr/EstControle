package main

import "fmt"

func main() {

	x:= 30
	if (x == 2 || x == 3) {
		fmt.Println("sim, x é 2 ou 3")
	}else{
		fmt.Println("não, x não é 2 ou 3")
	}

	if (x%2 == 0 && x%3 == 0) {
		fmt.Println("sim, x é multiplo de 2 ou 3")
	}else{
		fmt.Println("não, x não é multiplo de 2 ou 3")
	}

}