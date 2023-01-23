package main

import "fmt"

func main() {

	for horas := 0; horas < 1; horas++ {
		for minutos := 0; minutos < 60; minutos++ {
			for segundos := 0; segundos < 60; segundos++ {
				fmt.Printf("%02d:%02d:%02d\n", horas, minutos, segundos)
			}
		}
	}
}