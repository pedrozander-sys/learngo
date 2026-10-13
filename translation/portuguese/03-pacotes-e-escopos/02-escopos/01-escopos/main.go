// Copyright © 2018 Inanc Gumus
// Curso Aprenda Programação em Go
// Licença: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Mais tutoriais         : https://learngoprogramming.com
// Treinamento presencial : https://www.linkedin.com/in/inancgumus/
// Siga-me no twitter     : https://twitter.com/inancgumus

package main

// escopo de arquivo
import "fmt"

// escopo de pacote
const ok = true

// escopo de pacote
func main() { // o escopo de bloco começa

	var hello = "Hello"

	// hello e ok são visíveis aqui
	fmt.Println(hello, ok)

} // o escopo de bloco termina
