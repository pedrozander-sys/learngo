## Onde guardar os arquivos de código-fonte que pertencem a um pacote?
1. Cada arquivo deve ir para um diretório diferente
2. Em um único diretório *CORRETA*


## Por que se usa uma cláusula package em um arquivo de código-fonte Go?
1. Para importar um pacote
2. Para informar ao Go que o arquivo pertence a um pacote *CORRETA*
3. Para declarar uma nova função

> **1:** Quem faz isso é a declaração `import`.
>
>
> **3:** Quem faz isso é a declaração `func`.
>
>


## Onde você deve colocar a `cláusula package` em um arquivo de código-fonte Go?
1. Como o primeiro código do arquivo *CORRETA*
2. Como o último código do arquivo
3. Pode colocar em qualquer lugar


## Quantas vezes você pode usar a `cláusula package` em um único arquivo de código-fonte?
1. Uma vez *CORRETA*
2. Nenhuma
3. Várias vezes


## Qual destas é um uso correto da `cláusula package`?
1. `my package`
2. `package main`
3. `pkg main`


## Qual destas está correta?
1. Arquivos do mesmo pacote não podem chamar as funções uns dos outros
2. Arquivos do mesmo pacote podem chamar as funções uns dos outros *CORRETA*


## Como executar vários arquivos Go?
1. go run *.*go
2. go build *go
3. go run go
4. go run *.go *CORRETA*

> **4:** Você também pode chamar assim (supondo que file1.go, file2.go e file3.go estejam no mesmo diretório): go run file1.go file2.go file3.go
>
>
