package main

// Não faço a minima ideia do que ta acontencendo aqui
import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func tarefa(nome string) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		time.Sleep(1 + time.Second) // atrasa o proximo comando em 1 segundo
		fmt.Printf("Tarefa %s...\n", nome)
	}
	fmt.Printf("Tarefa %s <Concluido>\n", nome)
}

func normal() {
	time.Sleep(1 + time.Second)
	fmt.Println("********** °.° **********")
}

func main() {
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		//utilizando goroutine
		go tarefa(strconv.Itoa(i)) // Itoa converte um inteiro para string

	}

	normal()

	wg.Wait()
	fmt.Println("Todas as tarefas foram Concluidas")
}
