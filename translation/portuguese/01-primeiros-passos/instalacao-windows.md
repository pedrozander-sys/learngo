# GUIA RÁPIDO DE INSTALAÇÃO NO WINDOWS

## OBSERVAÇÃO

Se você tem o gerenciador de pacotes [chocolatey.org](https://chocolatey.org/), pode instalar o Go facilmente assim:

```
choco install golang
```

## 1- Instale o editor Visual Studio Code

1. Instale, mas não abra ainda.
2. Acesse: [https://code.visualstudio.com](https://code.visualstudio.com)
3. Selecione Windows e comece o download
4. Execute o instalador

## 2- Instale o Git

1. Baixe e execute o instalador. Acesse: [https://gitforwindows.org](https://gitforwindows.org)
2. Selecione o VSCode como editor padrão
3. Marque todas as caixas de seleção
4. Selecione: _"Use Git from the Windows Command Prompt"_
5. Codificação (encodings): selecione a opção _"Checkout as is..."_.

## 3- Instale o Go

1. Acesse [https://golang.org/dl](https://golang.org/dl)
2. Selecione Windows e faça o download
3. Inicie o instalador

## 4- Configure o VS Code

1. Abra o VS Code; na aba de extensões à esquerda, procure por "go" e instale
2. Feche o VS Code completamente e abra de novo

3. Vá ao menu View e selecione **Command Palette**
    1. Ou simplesmente pressione `ctrl+shift+p`
    2. Digite: `go install`
    3. Selecione _"Go: Install/Update Tools"_
    4. Marque todas as caixas de seleção

## 5- Usando o Git-Bash

* Neste curso vou usar comandos bash. O bash é apenas uma interface de linha de comando usada no OS X e no Linux, e uma das mais populares. Então, se quiser usá-lo também, em vez da linha de comando padrão do Windows, você pode usar o git bash que instalou. Com o git bash, você digita comandos como se estivesse no OS X ou no Linux.

* Se não quiser usar o git bash, tudo bem também. A escolha é sua. Mas saiba que vou usar principalmente comandos bash, porque eles permitem comandos mais avançados.

* Você também pode preferir uma alternativa mais poderosa ao git bash: o [Subsistema do Windows para Linux (WSL)](https://docs.microsoft.com/en-us/windows/wsl/install-win10)

* **Para usar o git bash, siga estes passos:**
    1. Procure por git bash no menu Iniciar
    2. Ou, se houver, clique no ícone na sua área de trabalho

    3. Configure também o VS Code para usar o git-bash por padrão:
        1. Abra o VS Code
        2. Vá até a Command Palette
            1. Digite: `terminal`
            2. Selecione: _"Terminal: Select Default Shell"_
            3. E selecione: _"Git Bash"_

    4. **OBSERVAÇÃO:** Normalmente seus arquivos ficam em `c:\`; porém, no git bash, você os encontra no diretório `/c/`. É exatamente o mesmo diretório, só um atalho.

## É isso! Aproveite! 🤩

<div style="page-break-after: always;"></div>

> Mais tutoriais: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2018 Inanc Gumus
>
> Curso Aprenda Programação em Go
>
> [Clique aqui para ler a licença.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
