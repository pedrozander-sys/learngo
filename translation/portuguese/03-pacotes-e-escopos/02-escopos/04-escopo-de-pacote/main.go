// Copyright © 2018 Inanc Gumus
// Curso Aprenda Programação em Go
// Licença: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Mais tutoriais         : https://learngoprogramming.com
// Treinamento presencial : https://www.linkedin.com/in/inancgumus/
// Siga-me no twitter     : https://twitter.com/inancgumus

package main

import "fmt"

func main() {
	fmt.Println("Hello!")

	// os dois arquivos pertencem ao mesmo pacote
	// aqui chamamos o `bye()` de bye.go
	bye()
}

// EXERCÍCIO: Remova os comentários da função abaixo
//            e analise a mensagem de erro

// func bye() {
// 	fmt.Println("Bye!")
// }

// ***** EXPLICAÇÃO *****
//
// ERRO: a função "bye" foi "redeclared" (redeclarada)
//       em "this block" (neste bloco)
//
// "this block" significa = "pacote main"
//
// "redeclared" significa = você está usando o mesmo nome
//   no mesmo escopo de novo
//
// o escopo do pacote main é:
// todos os arquivos de código-fonte que estão no mesmo pacote main
