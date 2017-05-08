package main

import "fmt"

func main() {
	u := Usuario{
		"Dipper",
		"Didi",
		true,
	}

	a := Admin{
		Usuario{
			"Lain",
			"zero",
			false,
		},
		30,
	}
	MsgUsuario(u)
	MsgUsuario(a)
}

func MsgUsuario(u UserInterface) {
	fmt.Println(u.Mostre())
}

type UserInterface interface {
	Mostre() string
}

type Usuario struct {
	nome   string
	nick   string
	online bool
}

func (u Usuario) Mostre() string {
	//return "Olá, meu nome é " + u.nome + " e meu nickname é " + u.nick
	return fmt.Sprintf("Olá, meu nome é %s, e meu nickname é %s\n", u.nome, u.nick)
}

type Admin struct {
	Usuario
	idade int
}

func (a Admin) Mostre() string {
	//return "Olá, meu nome é " + a.nome + " e meu nickname é " + a.nick + ". Eu sou o Admin!"
	return fmt.Sprintf("Olá, meu nome é %s, e meu nickname é %s. Eu sou o Admin!\n", a.nome, a.nick)
}
