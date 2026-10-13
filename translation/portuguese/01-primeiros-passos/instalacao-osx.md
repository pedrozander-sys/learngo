# GUIA RÁPIDO DE INSTALAÇÃO NO OS X

## OBSERVAÇÃO

Se você tem o [homebrew](https://brew.sh) instalado, pode instalar o Go facilmente assim:

```
# se você não tem o git, instale assim:
brew install git

# depois instale o go
brew install go

# adicione o caminho do GOBIN ao seu PATH em ~/.bash_profile
export PATH=${HOME}/go/bin:$PATH
```

## 1- Instale o editor Visual Studio Code

1. Instale, mas não abra ainda.
2. Acesse: [https://code.visualstudio.com](https://code.visualstudio.com)
3. Selecione OS X (Mac) e comece o download
4. Descompacte o arquivo baixado e mova-o para o diretório `~/Applications`.

## 2- Instale o Git

1. Baixe e execute o instalador. Acesse: [https://git-scm.com/downloads](https://git-scm.com/downloads)
2. Selecione o VSCode como editor padrão

## 3- Instale o Go

1. Acesse [https://golang.org/dl](https://golang.org/dl)
2. Selecione OS X (Mac)
3. Inicie o instalador

## 4- Configure o VS Code

1. Abra o VS Code; na aba de extensões à esquerda, procure por "go" e instale
2. Feche o VS Code completamente e abra de novo

3. Vá ao menu View e selecione **Command Palette**
    1. Ou simplesmente pressione `cmd+shift+p`
    2. Digite: `go install`
    3. Selecione _"Go: Install/Update Tools"_
    4. Marque todas as caixas de seleção

4. Quando terminar, abra a **Command Palette** novamente
    1. Digite: `shell`
    2. Selecione: _"Install 'code' command in PATH"_
        1. **OBSERVAÇÃO:** Você não precisa fazer isso se estiver no Windows.

## É isso! Aproveite! 🤩

<div style="page-break-after: always;"></div>

> Mais tutoriais: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2018 Inanc Gumus
>
> Curso Aprenda Programação em Go
>
> [Clique aqui para ler a licença.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
