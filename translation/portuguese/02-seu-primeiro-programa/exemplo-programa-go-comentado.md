# Exemplo de programa Go comentado

```go
// package main é um pacote especial
// ele permite que o Go crie um arquivo executável
package main

/*
Este é um comentário de várias linhas.

a palavra-chave import torna outro pacote disponível
  para este "arquivo" .go.

import "fmt" permite acessar as funcionalidades do pacote fmt
  aqui neste arquivo.
*/
import "fmt"

// "func main" é especial.
//
// O Go precisa saber por onde começar.
//
// func main cria um ponto de partida para o Go
//
// Depois de compilar o código,
// o runtime do Go vai executar esta função primeiro
func main() {
	// depois de: import "fmt"
	// a função Println do pacote "fmt" fica disponível

	// Veja como ela é digitando no console:
	//   go doc -src fmt Println

	// Println é apenas uma função exportada do
	//   pacote "fmt"

	// Exportada = a primeira letra é maiúscula
	fmt.Println("Hello Gopher!")

	// O Go não chama a função Println sozinho.
	// Por isso você precisa chamá-la aqui.
	// Ele só chama `func main` automaticamente.

	// -----

	// O Go suporta caracteres Unicode em literais de string
	// E também no código-fonte: KÖSTEBEK!
	//
	// Porque: Literal ~= Código-fonte
}
```

<div style="page-break-after: always;"></div>

> Mais tutoriais: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2018 Inanc Gumus
>
> Curso Aprenda Programação em Go
>
> [Clique aqui para ler a licença.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
