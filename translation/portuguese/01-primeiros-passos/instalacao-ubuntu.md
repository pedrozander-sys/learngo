# Guia rápido de instalação no Linux

Se quiser usar o snap, você pode instalar o Go facilmente assim:

    sudo snap install go --classic
    
Caso contrário, siga os passos abaixo:

## 1. Atualize os pacotes locais

  ```bash
  sudo apt-get update
  ```

## 2. Instale o git

  ```bash
  sudo apt-get install git
  ```

## 3. Instale o Go

Há duas formas:

1- Pela web: selecione Linux e o download vai começar.

  ```bash
  firefox https://golang.org/dl
  ```

2- Usando o snap: se escolher esta opção, pule para o 5º passo.

  ```bash
  sudo snap install go --classic
  ```

## 4. Copie o Go para o diretório correto

1. Descubra o nome do arquivo baixado

2. Use esse nome para descompactá-lo

    ```bash
    gofile="DELETE_THIS_AND_TYPE_THE_NAME_OF_THE_DOWNLOADED_FILE_HERE (without its extension)"

    tar -C /usr/local -xzf ~/Downloads/$gofile
    ```

    _(Substitua o texto entre aspas pelo nome do arquivo baixado, sem a extensão.)_

## 5. Adicione o diretório de executáveis do Go ao seu PATH

1. Adicione o diretório `go/bin` ao `$PATH` para poder executar os comandos fundamentais do Go.

    ```bash
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.profile
    ```

2. Adicione o diretório "$HOME/go/bin" ao $PATH

    ```bash
    echo 'export PATH=$PATH:$HOME/go/bin' >> ~/.profile
    ```

## Instale as ferramentas do Go:

* São ferramentas muito úteis para facilitar o desenvolvimento (como o goimports)

* O `go get` não funciona sem um programa de controle de versão como o Git, que já instalamos acima.

* Isso vai criar o diretório `~/go` e baixar as ferramentas do Go para lá.

    * Esse diretório também é onde você deve colocar o seu código.
    (Se não for usar Go Modules)

    ```bash
    go get -v -u golang.org/x/tools/...
    ```

## Instale o VSCode (opcional)

Observação: você pode usar outro editor de código, se preferir. Porém, o curso usa o Visual Studio Code (VSCode).

1. Abra o aplicativo "Ubuntu Software"

2. Procure por VSCode e clique em "Install"


## PASSO OPCIONAL:

1. Crie um arquivo hello.go em um novo diretório, em qualquer lugar fora do `$GOPATH`

    ```bash
    cat <<EOF > hello.go
    package main

    import "fmt"

    func main() {
        fmt.Println("hello gopher!")
    }
    EOF
    ```

2. Execute o programa

    ```bash
    go run hello.go
    Deve imprimir: hello gopher!
    ```

<div style="page-break-after: always;"></div>

> Mais tutoriais: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2018 Inanc Gumus
>
> Curso Aprenda Programação em Go
>
> [Clique aqui para ler a licença.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
