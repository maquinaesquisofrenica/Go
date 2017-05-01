package main

import "fmt"

func main() {

	var objeto Lista

	objeto = objeto.Iniciar() // iniciando por retorno
	objeto.Mostre()

	var a Lista
	a.Iniciar_Ref() //inciando por referencia
	fmt.Println(a)
}

type Lista []interface{} // uma interface vazia aceita qualquer tipo

func (obj Lista) Mostre() { // este funçao so pode ser chamada por uma variavel do tipo criado acima
	fmt.Println(obj)
}

func (obj Lista) Iniciar() Lista {
	obj = []interface{}{
		3.14159,
		"Guitarra",
		1000,
		true,
	}
	return obj
}

func (obj *Lista) Iniciar_Ref() {
	*obj = []interface{}{
		3.14159,
		"Guitarra",
		1000,
		true,
	}
}
