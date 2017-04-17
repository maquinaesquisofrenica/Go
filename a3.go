package main

import ("fmt"
        "math")

func main(){
   // em go só exite uma estrutura de repetição que é o FOR
   var numero int = 0;
   for numero < 10 { // utilizando o FOR como se fosse um WHILE
     fmt.Println(numero)
     numero++;
   }

   //fmt.Println(math.Sqrt(float64(numero))) // math.Sqrt retorna a raiz do valor, porem so aceitao valores float64
   // float64(numero) converte o valor para float64

   for i := 10; i >= 0; i-- {
     fmt.Println("A raiz de,",i,"é", math.Sqrt(float64(i)))
   }
}
