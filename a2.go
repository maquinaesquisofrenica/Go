package main

import (
	"fmt"
	"math/rand" // contem funções de numeros randomicos
	"strconv"   // pacote que contem a função Itoa que converte inteiro para string
	"time"      // funções utilizando o relogio do sistema
)

func main() {
	b := true
	teste(b)
	var x int
	x = 10
	// em go não aceita parentes no if
	if x%2 == 0 {
		fmt.Println(strconv.Itoa(x) + " e Par")
	} else {
		fmt.Println(strconv.Itoa(x) + "e Impar")
	}


  if x:= 1; x > 0{  // declarando uma variavel dentro do if
	   fmt.Println(num,"é maior que 0") // variavel x e só pode ser utilizada dentro do escopo do if
		 // a variavel x que está no if e diferente da variavel x criada mais acima do codigo
  }

	rand.Seed(int64(time.Now().Nanosecond()))
	var number int
	number = rand.Int() % 3
	fmt.Println("Numero gerado", number)
	switch number {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("um")
	default:
		fmt.Println("dois")
	} //fim switch

}

func teste(param bool) {
	if param {
		fmt.Println("Variavel true")
	} else {
		fmt.Println("Variavel false")
	}
}
