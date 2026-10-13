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

	// Você pode acessar funções de outros arquivos
	// que estão no mesmo pacote

	// Aqui, `main()` pode acessar `bye()` e `hey()`

	// Isso porque bye.go, hey.go e main.go
	//   estão no pacote main.

	bye()
	hey()
}
