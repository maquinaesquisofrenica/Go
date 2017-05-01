package main

import "fmt"

func main() {
	a := livro{"O Chamdo de Cthulhu ", "H.P. Lovecraft", 300}
	fmt.Println(a.Titulo)
	fmt.Println(a.Autor)
	fmt.Println(a.Paginas)

	a.Biblioteca()

	//forma alternativa caso nao saiba qual a sequencia dos tipos
	b := livro{
		Paginas: 200,
		Autor:   "JK Rowling",
		Titulo:  "Harry Potter e o Cálice de Fogo",
	}

	fmt.Println(b)
	b.Descricao("Recomendado!")

}

type livro struct {
	Titulo  string
	Autor   string
	Paginas int
}

//criando um metodo que pertence ao tipo da struct criada
func (a *livro) Descricao(coment string) {
	fmt.Println(coment)
}
