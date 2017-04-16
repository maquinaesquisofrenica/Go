package main //pacote obrigatorio para a execução de um programa em go

/*
  Para executar um programa em go, acesse o local em que foi salvo
  este arquivo pelo terminal ou cmd(windows)
  Comandos para executar:

  go run nome_do_arquivo.go  <- Este comando executa o codigo e apaga o executavel em seguida

  go build nome_do_arquivo.go <- Este comando gera uma executavel do codigo

  ./nome_do_arquivo   <-Este comando executa o executavel

*/

import "fmt" // pacote que contem as funções Print

var num1, num2 int = 100, 10

func main() { // func main e a função obrigatorio para a execução de um programa em go

	fmt.Println("Golang - aula 01")
	fmt.Println("\nApresentação")
	fmt.Println()

	print("Olá Mundo!\n") // printf() e uma função que imprime uma string na tela

	//  \n = encerra a linha

	// Não é aconselhavel usar a função printf desta forma em go

	fmt.Println("Tudo ok!") // esta e a melhor forma de usar uma função Print

	// Variáveis

	fmt.Println("\nDeclaração de variáveis:")
	fmt.Println("")

	var name string
	name = "Lain"

	fmt.Printf("name: %s\n", name) // %s indica que esse local será substituido por uma string

	name = "Dipper" // sobreescrevendo a variavel

	fmt.Println("name: " + name) // concatenação de strings usando o operador +

	fmt.Print("Bem Vindos ao Go!\n\n")

	var linguagens, autor string
	linguagens, autor = "Golang", "Eric" // passando valores em varias variaveis inline

	fmt.Printf("Autor: %s, linguagem: %s\n\n", autor, linguagens)

	//Criando blocos de variaveis

	var (
		day    int
		semana string
		year   int32
	)

	day, semana, year = 17, "Domingo", 2017
	fmt.Printf("Date: %d/%s/%d\n\n", day, semana, year)

	// Declarações sem especificar Tipos

	var livro = "Necronomicom" // o compilador reconhece que é uma string
	// se a variavel acima receber valores de outro tipo isso irá causar um erro
	fmt.Println("livro: " + livro)
	fmt.Println()

	game := "Residente Evil 7" // Operador especial nao funciona fora de funçoes
	// operador especial (:= ) entende que na linha acima a uma variavel que está recebendo um valor

	fmt.Println("game: " + game)
	fmt.Println()

	age, peso := 25, 89.03 // inicializando multiplas variáveis com operador especial

	fmt.Printf("Idade: %d, Peso: %.2f\n\n", age, peso)

	// Tipos de Variáveis

	fmt.Println("Tipos de variáveis:")
	fmt.Println()

	fmt.Println("Tipo bool:")
	var boleano bool // variavel que recebe true ou false
	boleano = true
	fmt.Print("Valor da variavel \"boleano\": ")
	fmt.Println(boleano)
	fmt.Print("Valor de !boleano: ")
	fmt.Println(!(boleano))
	fmt.Println()
	fmt.Println("Tipos int:")
	var inteiro int // recebe numeros inteiros
	/*
	   variações de inteiros:
	    var inteiro int8
	    var inteiro int16
	    var inteiro int32
	    var inteiro int64
	*/
	inteiro = -54
	fmt.Printf("Valor da variavel inteiro: %d\n\n", inteiro) // %d indica que será substituido por um numero inteiro

	fmt.Println("Tipos float:")
	var real1 float32 // recebe numeros reais de até 32 bits
	var real2 float64 // recebe numeros reais de até 64 bits

	real1 = 1.618
	real2 = 3.14159

	fmt.Printf("valor da variavel real1: %f\n", real1)
	fmt.Printf("Valor da variável real2: %f\n", real2)

	fmt.Println("\nTipo byte:")
	var bait byte
	bait = 130
	fmt.Print("variavel bait: ")
	fmt.Print(bait)
	fmt.Println()

	fmt.Println("\nConstantes:")
	const dog = "cachorro" // constantes são inicializas uma vez e não podem mais ser alteradas
	fmt.Println("valor da constante dog: " + dog)

	//Constantes em blocos
	const (
		anime    = "Serial Experiments Lain"
		password = 123789
		money    = 317.55
	)
	fmt.Println("Bloco de constantes:")
	fmt.Printf("anime: %s\nsenha: %d\ndinheiro: %.2f R$\n\n", anime, password, money)

	fmt.Println("Operações Básicas:")
	fmt.Printf("%d + %d = %d\n", num1, num2, (num1 + num2))
	fmt.Printf("%d - %d = %d\n", num1, num2, (num1 - num2))
	fmt.Printf("%d * %d = %d\n", num1, num2, (num1 * num2))
	fmt.Printf("%d / %d = %d\n", num1, num2, (num1 / num2))
	fmt.Printf("resto da divisão de %d por %d = %d\n", num1, num2, (num1 % num2))

}
