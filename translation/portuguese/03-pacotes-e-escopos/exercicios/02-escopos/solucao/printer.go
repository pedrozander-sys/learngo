// Copyright © 2018 Inanc Gumus
// Curso Aprenda Programação em Go
// Licença: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Mais tutoriais         : https://learngoprogramming.com
// Treinamento presencial : https://www.linkedin.com/in/inancgumus/
// Siga-me no twitter     : https://twitter.com/inancgumus

package main

import "fmt"

func hello() {
	// só este arquivo pode acessar o pacote fmt importado
	// quando os outros também o importam, eles também podem acessar
	//   o seu próprio "nome" `fmt`

	fmt.Println("hi! this is inanc!")
	bye()
}
