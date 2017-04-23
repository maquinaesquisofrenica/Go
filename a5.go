package main

import "fmt"

func main() {

	fmt.Println("Arrays/Vetores:")
	var vetor1 [10]int // declarando um Array/Vetor com espaço para 10 inteiros
	// vetor := [10]int{}  <- está expressão também seria valida

	// atribuindo valor para um posicao do vetor
	vetor1[0] = 10
	vetor1[9] = 100

	// imprimindo todos os valores do vetor
	fmt.Println(vetor1)

	// len() retorna a quantidade de espços de um Array/Vetor
	fmt.Println("O tamanho deste vetor é:", len(vetor1))

	vetor2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Array/Vetor com valores inicializados
	fmt.Println(vetor2)
	fmt.Println("Valor da posicao 5 é", vetor2[5])

	//var nome string = "Wolverine"
	//fmt.Printf("%s tem %d letras\n", nome, len(nome)) // len() funciona com strings

	fmt.Println("\nMap:")
	// declarando um map
	pessoa1 := map[string]string{ // um map recebe o tipo do indice e o tipo da variavel dentro dele
		"nome":  "Laura", // a string com valor "Laura" é atribuida ao indice "nome"
		"idade": "13 anos",
		"poder": "fator de cura + adamatium",
	}

	//Atribuindo o indice e um valor no Map pessoa1
	//pessoa1["nome"] = "Laura"
	fmt.Println(pessoa1["nome"])
	fmt.Println(pessoa1["idade"])
	fmt.Println(pessoa1["poder"])

	// função len() em maps, retorna a quantidade de indices
	fmt.Println("Tamanho do map:", len(pessoa1))

	games := map[string]string{}
	games["jogo"] = "Super Metroid"
	games["console"] = "Super Nintendo"
	// cada indice é incluido automaticamente no map
	fmt.Println(games)
	games["lançamento"] = "1991"
	fmt.Println(games)

	fmt.Println("\nSlices:")
	// slice é um array que não tem limite pré definido
	slice := []int{10, 20, 30}
	//slice[3] = 40 <- este tipo de atribuição não funciona num slice

	//a função append insere valores no final do slice
	slice = append(slice, 40, 50, 60)
	fmt.Println(slice)

}
