package main

import "fmt"

func main() {
	var a Admin
	a.user.nome = "Lain"
	a.user.idade = 19
	a.user.online = true
	a.ranking = 1
	fmt.Println(a)

	var u simpleUser
	u.Menssagem()

	fmt.Println("\nHerança por Composição")
	var mulher Genero
	mulher.nome = "Mabel"
	mulher.idade = 25
	mulher.peso = 55.9
	mulher.sexo = "feminino"
	fmt.Println(mulher)
	mulher.Aniversario()

	//Inicializando na Declaraçao
	homem := Genero{
		Pessoa{
			"Dipper",
			26,
			61.4,
		},
		"masculino",
	}
	fmt.Println(homem)
	homem.Aniversario()
}

type Usuario struct {
	nome   string
	idade  int
	online bool
}

type simpleUser Usuario

func (us simpleUser) Menssagem() {
	fmt.Println("Hello friend")
}

type Admin struct {
	user    Usuario
	ranking int
}

// HERANÇA POR COMPOSIÇÃO

type Pessoa struct {
	nome  string
	idade int
	peso  float32
}

func (a *Pessoa) Aniversario() {
	a.idade++
	fmt.Println(a.nome, "Completou", a.idade, "anos")
}

type Genero struct {
	Pessoa // a struct Genero herda tudo  da struct Pessoa
	sexo   string
}
