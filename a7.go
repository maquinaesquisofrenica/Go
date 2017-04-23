package main

import "fmt"

func main() {
	vetor := [...]int32{1, 2, 3}
	// nesta sintaxe o tamanho do vetor será definido pela quantidade de valores
	// que for atribuido a ele na declaração
	fmt.Println("Loop (for) padrao:")
	for i := 0; i < len(vetor); i++ {
		fmt.Println(vetor[i])
	}

	fmt.Println("Loop (for) range:")
	//percorre todo o vetor sem precisar informar o indice de parada
	for indice, valor := range vetor {
		fmt.Println("indice:", indice, "=", valor)
	}

	fmt.Println("\nLoop (for) range com Maps:")
	mutante := map[string]string{
		"nome":     "Laura",
		"codenome": "x-23",
		"idade":    "15 anos",
	}
	// o for irá percorrer todos os indices do map
	for chave, valor := range mutante {
		fmt.Println(chave, ":", valor)
	}

	fmt.Println("\nLoop (for) range com Slices:")
	medias := []float32{5.5, 7.8, 10.0, 9.2, 3.75}
	medias = append(medias, 8.8, 7.5)
	for indice, valor := range medias {
		fmt.Println("indice:", indice, "media:", valor)
	}

	//exemplo com for range ignorando o indice
	for _, valor := range medias {
		fmt.Println(valor)
	}
}
