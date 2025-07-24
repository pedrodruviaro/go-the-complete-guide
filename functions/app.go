package main

import "fmt"

func main() {
	printSomething("Functions!")

	s, m := operations(3, 5)
	printSomething(s)
	printSomething(m)
}

func printSomething(value any) {
	fmt.Println(value)
}

// ou -> a int, b int
func operations(a, b int) (int, int) {
	sum := a + b
	multiply := a * b

	return sum, multiply
}

// podemos criar variáveis automaticamente com base no tipo do retorno
func operations2(a, b int) (sum int, multiply int) {
	sum = a + b
	multiply = a * b

	return
}