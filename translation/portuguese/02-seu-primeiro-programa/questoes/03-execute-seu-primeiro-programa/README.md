## Qual é a diferença entre `go build` e `go run`?
1. `go run` apenas compila um programa; já `go build` compila e executa.
2. `go run` compila e executa um programa; já `go build` apenas compila. *CORRETA*

> **1:** Na verdade, é o contrário.
>
>
> **2:** `go run` compila o seu programa e o coloca em um diretório temporário. Depois executa o programa compilado de lá.
>
>


## O Go salva o código compilado em um diretório. Qual é o nome desse diretório?
1. O mesmo diretório onde você chama `go build` *CORRETA*
2. O diretório $GOPATH/src
3. O diretório $GOPATH/pkg
4. Um diretório temporário.

> **2:** Ali ficam apenas arquivos de código-fonte Go
>
>
> **3:** O Go só coloca o seu código ali quando você chama `go install`.
>
>


## O que é verdade sobre o runtime (tempo de execução)?
1. Acontece quando o seu programa começa a rodar em um computador *CORRETA*
2. Acontece enquanto o seu programa está sendo compilado


## O que é verdade sobre o compile-time (tempo de compilação)?
1. Acontece quando o seu programa começa a rodar em um computador
2. Acontece enquanto o seu programa está sendo compilado  *CORRETA*


## Quando um programa Go pode imprimir uma mensagem no console?
1. Enquanto está sendo compilado.
2. Enquanto está rodando (depois do tempo de compilação). *CORRETA*
3. Enquanto está rodando (dentro do tempo de compilação).

> **1:** Na etapa de compilação, o seu programa não pode imprimir mensagens. Nesse estágio, ele está literalmente morto.
>
>
> **2:** Isso mesmo. É o único momento em que o seu programa pode interagir com um computador e mandar imprimir uma mensagem no console.
>
>
> **3:** A execução só pode acontecer depois do tempo de compilação
>
>
