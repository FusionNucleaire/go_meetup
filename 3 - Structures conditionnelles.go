package main

import  (
	"fmt"
)

func main() {

	fmt.Print("Nombres de 1 à 20\n")

	for i := 1; i <= 20; i++ {
		if i % 2 == 0 {
			fmt.Printf("Nombre pair : %d\n", i)
		} else {
			fmt.Printf("Nombre impair : %d\n", i)
		}
	}

}