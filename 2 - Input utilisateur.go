package main

import  (
	"fmt"
	"bufio"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entrez du texte : ")
	input, _ := reader.ReadString('\n')
	fmt.Print("Ce que vous avez entré : ", input)
}