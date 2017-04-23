package main

import "fmt"

func main() {
	funcao()
}

func funcao() {
	fmt.Println("Olá, eu sou uma função")
	soma(10, 10)
	fmt.Println("Função retornou: ", getSoma(3.14159, 2.54))

	var x, y float64 = getValores(getSoma(10, 1), getSoma(10, 2)) // esta função retorna dois  valores
	fmt.Println("valor de x = ", x)
	fmt.Println("valor de y = ", y)

	_, z := getValores(1, 2)
	// o primeiro valor retornado que sera atribuido pelo underline será descartado
	fmt.Println("valor de z = ", z)

	var a, b int
	fmt.Print("\nNumerador: ")
	fmt.Scan(&a)
	fmt.Print("Denominador: ")
	fmt.Scan(&b)
	div, e := divisao(a, b)
	if e != "" {
		fmt.Println(e)
	} else {
		fmt.Printf("%d / %d = %d\n", a, b, div)
	}
}

func soma(a, b int64) { // caso os tipo sejam iguais pode-se declara-lo apenas uma vez
	fmt.Printf("%d + %d = %d\n", a, b, (a + b))
}

func getSoma(a, b float64) float64 { // float64 fora do parenteses é o tipo do retorno da função
	return (a + b)
}

func getValores(a, b float64) (float64, float64) { // função que retorna dois valores ao mesmo tempo
	return a, b
}

func divisao(a, b int) (int, string) {
	if b == 0 {
		return 0, "divisao por 0"
	}
	return (a / b), ""
}
