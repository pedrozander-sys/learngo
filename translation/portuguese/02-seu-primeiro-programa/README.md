# Guia rápido: Escreva seu primeiro programa em Go

Olá!

Você pode guardar este guia rápido para consulta depois de assistir às aulas desta seção.

Você também pode imprimi-lo e acompanhar as videoaulas desta seção com ele.

Aproveite!

---

## COMANDOS DE LINHA DE COMANDO:

* Entrar em um diretório: `cd caminhoDoDiretorio`

* **WINDOWS:**

    * Listar os arquivos de um diretório: `dir`

* **OS X E LINUX:**

    * Listar os arquivos de um diretório: `ls`

## COMPILANDO E EXECUTANDO PROGRAMAS GO:

* **Compilar um programa Go:**

    * Dentro do diretório do programa, digite:
        * `go build main.go`

* **Executar um programa Go:**

    * Dentro do diretório do programa, digite:
        * `go run main.go`

## ONDE VOCÊ DEVE COLOCAR SEUS ARQUIVOS-FONTE?

* Em qualquer diretório que você quiser.

## PRIMEIRO PROGRAMA

### Crie um diretório
* Crie um novo diretório:
  * `mkdir nomeDoMeuDiretorio`
* Entre nesse diretório:
  * `cd nomeDoMeuDiretorio`

### Adicione os arquivos de código-fonte
* Crie um novo arquivo com `code main.go`.
  * Isso vai criar o arquivo na pasta e abri-lo no Visual Studio Code.
* Adicione o código abaixo e salve.

```go
package main

import "fmt"

func main() {
    fmt.Println("Hi! I want to be a Gopher!")
}
```

### Execute o programa
* Por fim, volte para a linha de comando.
  * Execute assim: `go run main.go`
* Se você criar outros arquivos e quiser executar todos, use este comando:
  * `go run .`

É isso! Aproveite!

> Mais tutoriais: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2018 Inanc Gumus
>
> Curso Aprenda Programação em Go
>
> [Clique aqui para ler a licença.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
