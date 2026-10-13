## Qual palavra-chave abaixo você precisa usar para definir um pacote?
```go
package main

func main() {
}
```
1. func
2. package *CORRETA*
3. fmt.Println
4. import

> **1:** A palavra-chave func é usada para declarar uma nova função.
>
>
> **2:** Isso mesmo! A palavra-chave package permite definir um pacote para um arquivo Go.
>
>
> **3:** Não é uma palavra-chave, é uma função do pacote fmt.
>
>
> **4:** A palavra-chave import é usada para importar um pacote.
>
>


## Qual é o objetivo de usar package main no programa abaixo?
```go
package main

func main() {
}
```
* Criar um pacote de biblioteca
* Sair do programa corretamente
* Criar um programa Go executável *CORRETA*


## Qual é o objetivo de func main no programa abaixo?
```go
package main

func main() {
}
```
1. Define um pacote chamado main
2. Permite que o Go comece a executar o programa *CORRETA*
3. Imprime uma mensagem no console

> **1:** A função main não cria um pacote.
>
>
> **2:** Isso mesmo. O Go chama a função main automaticamente para executar um programa.
>
>
> **3:** Ela não imprime nada (pelo menos não diretamente).
>
>


## Qual é o objetivo de import "fmt" no programa abaixo?
```go
package main
import "fmt"

func main() {
    fmt.Println("Hi!")
}
```
1. Imprime "fmt" no console
2. Define um novo pacote chamado "fmt"
3. Importa o pacote `fmt`, para que você possa usar as funcionalidades dele *CORRETA*

> **1:** Quem imprime a mensagem é `fmt.Println`, não `import "fmt"`.
>
>
> **2:** Quem faz isso é a palavra-chave `package`, não a `import`.
>
>
> **3:** Sim. Por exemplo, depois de importar o pacote fmt, você pode chamar a função Println dele para imprimir uma mensagem no console.
>
>


## Qual palavra-chave é usada para declarar uma nova função?
* func *CORRETA*
* package
* Println
* import


## O que é uma função?
1. É como um miniprograma. É um bloco de código reutilizável e executável. *CORRETA*
2. Permite que o Go execute um programa.
3. Permite que o Go importe um pacote chamado function.
4. Imprime uma mensagem no console.

> **2:** Para isso, o Go procura package main e func main. Uma função sozinha não faz isso.
>
>
> **3:** Quem faz isso é a palavra-chave `import`.
>
>
> **4:** Por exemplo: `fmt.Println` faz isso.
>
>


## Você precisa chamar a função main por conta própria?
1. Sim, para que eu possa executar o meu programa.
2. Não, o Go chama a função main automaticamente. *CORRETA*

> **1:** Não, você não precisa chamar a função main. O Go a executa automaticamente.
>
>


## Você precisa chamar uma função para executá-la?
_(exceto a func main)_
1. Sim, para que o Go possa executar essa função. *CORRETA*
2. Sim, para que o Go possa executar o meu programa.
3. Não, o Go chama as funções automaticamente.

> **1:** Isso mesmo. Você mesmo precisa chamar uma função. O Go não vai executá-la automaticamente. O Go só chama automaticamente a função main (e algumas outras funções que você ainda não aprendeu).
>
>

> **2:** Esse é o papel apenas da `func main`. Só existe uma `func main`.
>
>

> **3:** O Go não chama nenhuma função automaticamente, exceto a func main (e algumas outras funções que você ainda não aprendeu). Então, exceto a func main, você precisa chamar as funções por conta própria.
>


## O que o programa abaixo imprime?
```go
package main

func main() {
}
```
1. Imprime uma mensagem no console
2. É um programa correto, mas não imprime nada *CORRETA*
3. É um programa incorreto

> **1:** Ele não imprime nenhuma mensagem. Para isso, você pode usar a função fmt.Println.
>
>

> **2:** Sim, é um programa correto; porém, como não contém fmt.Println, não imprime nada.
>
>

> **3:** É um programa correto. Ele usa a palavra-chave package e tem uma função main. Então é um programa Go válido e executável.
>
>


## O que este programa imprime?
```go
package main

func main() {
    fmt.Println(Hi! I want to be a Gopher!)
}
```
* Hi! I want to be a Gopher!
* Não imprime nada
* Este programa está incorreto *CORRETA*

> **1:** Ele não passa a mensagem para o Println entre aspas duplas. Deveria ser assim: fmt.Println("Hi! I want to be a Gopher")
>
>

> **3:** Ele não importa o pacote "fmt". Veja também o item 1.
>
>


## O que este programa imprime?
```go
package main
import "fmt"

func main() {
    fmt.Println("Hi there!")
}
```
* Hi there! *CORRETA*
* fmt
* Este programa está incorreto; ele importa o pacote errado ou não existe uma função chamada `Println`

> **2:** import "fmt" importa o pacote `fmt`, para que você possa usar as funcionalidades dele.
>
>
> **3:** Na verdade, este programa está correto.
>
>
