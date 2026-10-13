1. **Imprima o seu nome e o nome do seu melhor amigo** usando Println duas vezes. [Veja este exercício aqui](01-imprimir-nomes).

2. **Diga olá para você mesmo.**

    1. Compile o seu programa usando `go build`

    2. **Envie para um amigo**

       Ele ou ela deve usar o mesmo sistema operacional que você.

       Por exemplo, se você usa Windows, o sistema da outra pessoa também deve ser Windows.

    3. **Envie o seu programa para um amigo com outro sistema operacional**

       Para isso, você deve compilar o programa para o sistema operacional dele.

       **Crie um executável para OSX:**
       `GOOS=darwin GOARCH=386 go build`

       **Crie um executável para Windows:**
       `GOOS=windows GOARCH=386 go build`

       **Crie um executável para Linux:**
       `GOOS=linux GOARCH=arm GOARM=7 go build`

       **A lista completa está aqui:**
       https://golang.org/doc/install/source#environment

       **OBSERVAÇÃO:** Se você usa o prompt de comando ou o PowerShell, talvez precise usar assim:
       `cmd /c "set GOOS=darwin GOARCH=386 && go build"`

3. **Chame [Print](https://golang.org/pkg/fmt/#Print) em vez de [Println](https://golang.org/pkg/fmt/#Println)** e veja o que acontece.

4. **Chame [Println](https://golang.org/pkg/fmt/#Println) ou [Print](https://golang.org/pkg/fmt/#Print) com vários valores**, separando-os por vírgulas.

5. **Remova as aspas duplas de um literal de string** e veja o que acontece.

6. **Mova as declarações package e import** para o fim do arquivo e veja o que acontece.

7. **[Leia a documentação online do Go](https://golang.org/pkg)**.

    1. Dê uma olhada rápida nos pacotes e leia o que eles fazem.

    2. Veja o código-fonte deles clicando nos títulos.

    3. Você não precisa entender tudo, só faça. Isso vai te preparar para as próximas aulas.

8. Faça também **o tour em**: https://tour.golang.org/

    1. Dê uma olhada rápida nos recursos da linguagem.

    2. Vamos falar sobre todos eles em breve.

9. [Siga-me no twitter para aprender mais](https://twitter.com/inancgumus).
