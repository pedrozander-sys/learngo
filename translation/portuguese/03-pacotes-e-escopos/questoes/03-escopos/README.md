## O que é um escopo?
* Um bloco de código executável
* A visibilidade dos nomes declarados **CORRETA**
* Determina o que executar

```go
package awesome

import "fmt"

var enabled bool

func block() {
    var counter int
    fmt.Println(counter)
}
```

## Qual nome abaixo tem escopo de pacote?
1. awesome
2. fmt
3. enabled **CORRETA**
4. counter

> **3:** Isso mesmo. `enabled` está fora de qualquer função, então é um nome com escopo de pacote. A função `block()` também tem escopo de pacote; ela também está fora de qualquer função.
>
>


## Qual nome abaixo tem escopo de arquivo?
1. awesome
2. fmt **CORRETA**
3. enabled
4. block()
5. counter

> **2:** Isso mesmo. Os nomes de pacotes importados têm escopo de arquivo e só podem ser usados dentro do mesmo arquivo.
>
>


## Qual nome abaixo está no escopo da func block()?
1. awesome
2. fmt
3. enabled
4. block()
5. counter **CORRETA**

> **5:** Isso mesmo. `counter` é declarado dentro da func `block()`, então está no escopo dela. Fora da func `block()`, nenhum outro código o enxerga.
>
>


## `block()` enxerga o nome `enabled`?
1. Sim: ele está no escopo de pacote **CORRETA**
2. Não: ele está no escopo de arquivo
3. Não: ele está no escopo de bloco de block()

> **1:** Todo código dentro do mesmo pacote enxerga todos os outros nomes declarados no nível do pacote.
>
>


## Outros arquivos do pacote `awesome` enxergam o nome `counter`?
1. Sim
2. Não: ele está no escopo de pacote
3. Não: ele está no escopo de arquivo
4. Não: ele está no escopo de bloco de block() **CORRETA**

> **4:** Isso mesmo. Nenhum outro código enxerga os nomes dentro da função `block()`. Só o código dentro de `block()` os enxerga (e só até certo ponto: dentro do bloco, o código só enxerga as variáveis declaradas antes dele).
>
>


## Outros arquivos do pacote `awesome` enxergam o nome `fmt`?
1. Sim
2. Não: ele está no escopo de pacote
3. Não: ele está no escopo de arquivo **CORRETA**
4. Não: ele está no escopo de bloco de block()

> **3:** Só o próprio arquivo enxerga os pacotes importados; os outros arquivos não, estejam no mesmo pacote ou não.
>
>


## O que acontece se você declarar o mesmo nome no mesmo escopo, assim:
```go
package awesome

import "fmt"

// declarado duas vezes no escopo de pacote
var enabled bool
var enabled bool

func block() {
    var counter int
    fmt.Println(counter)
}
```
1. O nome recém-declarado vai sobrescrever o anterior.
2. Não posso fazer isso. Ele já foi declarado no escopo de pacote. *CORRETA*
3. Não posso fazer isso. Ele já foi declarado no escopo de arquivo.

> **2:** Isso mesmo. Você não pode declarar o mesmo nome no mesmo escopo. Se pudesse, como acessaria esse nome depois? E qual dos dois?
>
>


## O que acontece se você declarar o mesmo nome em outro escopo, assim:
```go
package awesome

import "fmt"

// declarado no escopo de pacote
var enabled bool

func block() {
    // também declarado no escopo de bloco
    var enabled bool

    var counter int
    fmt.Println(counter)
}
```
1. O nome recém-declarado vai sobrescrever o anterior. *CORRETA*
2. Não posso fazer isso. Ele já foi declarado no escopo de pacote.
3. Não posso fazer isso. Ele já foi declarado no escopo de arquivo.

> **1:** Na verdade, você pode declarar o mesmo nome em escopos internos assim. O escopo de `block()` fica dentro do seu pacote. Isso significa que ele acessa o escopo do pacote (mas não o contrário). Ou seja, o escopo de `block()` está abaixo do escopo do pacote, então você pode declarar o mesmo nome de novo. Ele vai sobrescrever (sombrear) o nome do escopo pai. Os dois podem existir ao mesmo tempo. Veja o exemplo no repositório do curso para conferir.
>
>
