package main

import "fmt"

func main() {
	fmt.Println(soma(10, 20, 30, 1000))
	fmt.Println(soma()) // também funciona caso nao recebe nada, logicamente retornará 0

	fmt.Println("\nFunção recursiva")
	fmt.Println("\nSequencia de Fibonacci")
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacci(i))
	}

	fmt.Println("\nFatorial")
	var num int
	fmt.Print("Digite um numero: ")
	fmt.Scan(&num)
	fmt.Println("Fatorial de", num, " é ", fatorial(num))
}

//função variadica recebe um numero indeterminado de parametros
func soma(numeros ...int) int { //os parametros serão alocados em um slice
	resultado := 0
	for _, valor := range numeros {
		resultado += valor
	}
	return resultado
}

func fibonacci(num int) int {
	if num == 1 || num == 0 {
		return 1
	} else {
		return fibonacci(num-1) + fibonacci(num-2)
	}
}

func fatorial(num int) int {
	if num == 1 || num == 0 {
		return 1
	} else {
		return num * fatorial(num-1)
	}
}
