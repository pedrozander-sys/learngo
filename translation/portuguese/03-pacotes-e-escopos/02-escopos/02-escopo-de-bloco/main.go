// Copyright © 2018 Inanc Gumus
// Curso Aprenda Programação em Go
// Licença: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Mais tutoriais         : https://learngoprogramming.com
// Treinamento presencial : https://www.linkedin.com/in/inancgumus/
// Siga-me no twitter     : https://twitter.com/inancgumus

package main

func nope() { // o escopo de bloco começa

	// hello e ok só são visíveis aqui
	const ok = true
	var hello = "Hello"

	_ = hello
} // o escopo de bloco termina

func main() { // o escopo de bloco começa

	// hello e ok não são visíveis aqui

	// ERRO:
	// fmt.Println(hello, ok)

} // o escopo de bloco termina
