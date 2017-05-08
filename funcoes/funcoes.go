package funcoes

func Ola() string { // funções com a primeira letra maiuscula são public
	return "Ola mundo!"
}

func Welcome(nome, sexo string) string {
	if sexo == "F" {
		return "Seja Bem vinda, " + nome
	}
	return "Seja Bem vindo, " + nome
}

func DigaOi(nome string) string {
	return "Ola, " + nome
}

func privado() string { // funções com a primeira letra minuscula são private
	return "Função privada. So pode ser usada neste pacote"
}
