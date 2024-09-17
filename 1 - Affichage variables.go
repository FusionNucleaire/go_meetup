package main

import (
	"fmt"
)

func main() {
	var chaine string = "Contextuel"
	fmt.Println("Ceci est une chaîne : "+chaine)

	fmt.Println("\n")

	fmt.Println("Liste des caractères ?")
	for i := 0; i < len(chaine); i++ {
        fmt.Println(chaine[i])
    }

	fmt.Println("\n")

	fmt.Println("En fait ce sont des runes")
	for i := 0; i < len(chaine); i++ {
        fmt.Printf("%x ", chaine[i])
    }

	fmt.Println("\n")

	fmt.Println("VRAIE Liste des caractères ?")
	for pos, char := range chaine {
		fmt.Printf("Lettre %c commence à la position %d\n", char, pos)
	}

}