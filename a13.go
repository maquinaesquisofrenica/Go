package main

import (
	// caminho baseado no diretorio src que foi configurado nas variaveis de ambiente
	"aulas/funcoes"
	"aulas/funcoes/matematica"
	"fmt"
)

func main() {
	fmt.Println(funcoes.Ola())
	fmt.Println(funcoes.Welcome("Laura", "F"))
	fmt.Println(funcoes.DigaOi("Logan"))

	d := matematica.Delta(2, -4, 0)
	var x1, x2 float64 = matematica.Raizes(2, -4, 0)
	fmt.Println(matematica.Aquivo())
	fmt.Println("Delta: ", d, "raiz do delta", matematica.MySqrt(d))
	fmt.Println("x1 =", x1, "| x2 =", x2)
}
