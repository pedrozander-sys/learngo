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
	// como você pode ver, não preciso importar um pacote
	// e consigo chamar a função `hello` aqui.
	//
	// isso porque os nomes com escopo de pacote
	// são compartilhados dentro do mesmo pacote
	hello()

	// mas aqui não consigo acessar o pacote fmt sem
	// importá-lo.
	//
	// isso porque ele está no escopo de arquivo de printer.go,
	// que é quem o importa.

	// esta func main também pode chamar a função bye aqui
	// bye()
}

// printer.go pode chamar esta função
//
// isso porque agora a função bye está no escopo de pacote
// do pacote main.
//
// a func main também pode chamá-la.
func bye() {
	fmt.Println("bye bye")
}
