// Copyright © 2018 Inanc Gumus
// Curso Aprenda Programação em Go
// Licença: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Mais tutoriais         : https://learngoprogramming.com
// Treinamento presencial : https://www.linkedin.com/in/inancgumus/
// Siga-me no twitter     : https://twitter.com/inancgumus

package main

import "fmt"

// Não falei sobre isso em nenhuma aula
// Como observação, quis colocar aqui
// Por favor, dê uma olhada.

var declareMeAgain = 10

func nested() { // o escopo de bloco começa

	// declara a mesma variável
	// as duas podem existir ao mesmo tempo
	// esta pertence apenas a este escopo
	// a variável do pacote continua intacta
	var declareMeAgain = 5
	fmt.Println("inside nested:", declareMeAgain)

} // o escopo de bloco termina

func main() { // o escopo de bloco começa

	fmt.Println("inside main:", declareMeAgain)

	nested()

	// a declareMeAgain do nível do pacote não é afetada
	// pela mudança dentro da func nested
	fmt.Println("inside main:", declareMeAgain)

} // o escopo de bloco termina
