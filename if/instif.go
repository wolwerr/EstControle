package main

import "fmt"

func main() {
for i:=1; i<=15; i++ {
	if i% 2 == 0 {
		fmt.Println("Par")
	}else{
		fmt.Println("Impar")
	}
	fmt.Println(i)	
}
}

