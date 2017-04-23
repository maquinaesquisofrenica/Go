package main

import "fmt"

func main() {
	fmt.Println("Slices")
	// a função make recebe um tipo, uma quantidade e uma capacidade
	// a quantidade e o numero de espaços inicializados
	// a capacidade é o numero de espaços limite
	exe1 := make([]string, 2, 3)
	//exe1 é dado do tipo slice que recebe strings
	//exe aponta para um array com 2 espaços mas que tem a capacidade de 3 espaços

	exe1[0] = "Laura"
	exe1[1] = "Logan"
	//exe[2] = "Xavier" <- não é possível adicionar
	exe1 = append(exe1, "Xavier") // adiciona o novo valor
	// o slice exe continua apontando para o mesmo endereço
	// pois o endereço ainda tinha capacidade para mais um valor
	fmt.Println(exe1)

	exe2 := exe1[:1] // corta todos os valore a partir do indice 1 até o fim
	fmt.Println("cortando o ultimo valor", exe2)
	exe3 := exe1[1:] // corta todos os valores a partir do indice (N-1) até o inicio
	fmt.Println("cortando o primeiro valor", exe3)
	exe4 := exe1[:] // recebe todos os valores do slice exe1
	fmt.Println("todos os valores:", exe4)

	//o comando abaixo estoura a capacidade que o array que o slice aponta tinha
	//exe1 = append(exe1, "Mistica") // este slice não aponta mais para o mesmo endereço
	//significa que os outros slices não tem mais a referencia de exe1
	// e eles nao podem mais modifica-lo

	exe2[0] = "Mistica"
	exe3[0] = "Jean Grey"
	//os valores do slice exe1 foram modificados pois um slice é uma referencia
	// ou seja, quando ele for atribuido a outro slice ele passa seu proprio endereço de memoria
	// é como se exe2 fosse um ponteiro
	fmt.Println("Valores modificados por outros slices:")
	fmt.Println(exe1)
	fmt.Println("Tamanho do slice:", len(exe1))

}
