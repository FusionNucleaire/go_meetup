package main

import  (
	"fmt"
	"strconv"
)

func fib(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	var n1 int = 0
	var n2 int = 1
	var somme int = 0
	for i := 2; i <= n; i++ {
		somme = n1 + n2
		n1 = n2
		n2 = somme
	}

	return n2		
}

func main() {

	fmt.Print("Entrez un nombre : ")

	var text1 string
	fmt.Scanln(&text1)

	fmt.Printf("Type : %T, Nombre entré : %v\n", text1, text1)

	int1, _ := strconv.Atoi(text1)

	fmt.Printf("Résultat de Fibonacci pour %d : %d", int1, fib(int1))

}