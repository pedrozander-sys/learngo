# ROTEIRO DO PROGRAMADOR

Olá!

Se você é um desenvolvedor experiente, talvez queira assistir apenas às aulas abaixo. Seguindo este roteiro, você vai precisar assistir a cerca de 50 aulas em vez de 180.

Este curso começa do mais básico e avança, passo a passo, até o fim. Ou seja, a complexidade dos assuntos aumenta a cada nível. Projetei assim de propósito, para facilitar para todo mundo.

Se achar que alguns assuntos são fáceis para você, assista às aulas de revisão (recap), faça os questionários e exercícios e até pule as aulas daquela seção inteira; você sempre pode voltar a elas depois.

Aproveite!

## AULAS

_(Os títulos das aulas foram mantidos em inglês, como aparecem no curso.)_


* **Write Your First Go Program** (Escreva seu primeiro programa em Go)
  * Assista a todas as aulas.
  * What is Go Doc?
  * As aulas em "Write a Library Package".

* **Master the Type System of Go** (Domine o sistema de tipos do Go)
  * Every Go type has a zero value
  * What is a blank identifier?
  * Let's declare a couple of variables!
  * What is type inference?
  * How to short declare multiple variables?
  * Why can't you short declare a variable in the package-level?
  * What is redeclaration?
  * When to use a short declaration?
  * Get input from command-line and learn about slices
  * Learn the basics of os.Args
  * Greet people using os.Args
  * As aulas em "Print Formatted Output Using Printf".
  * Convert Celsius to Fahrenheit
  * Convert Feet to Meters
  * What is a Raw String Literal?
  * How to get the length of a string?
  * As aulas a partir de "What is a Predeclared Type?" até o fim da seção.
  * As seguintes aulas em "Understand Untyped Constants"
    * Learn the rules of constants
    * Recap: Constants
    * How untyped constants work under the hood?
    * What is a Default Type?
    * Example: time.Duration
    * What is iota?
    * Naming Things: Recommendations

* **Control Flow and Error Handling** (Controle de fluxo e tratamento de erros)
  * Assista a todas as aulas em "Pass Me: Create a Password-Protected Program"
  * Assista a todas as aulas em "Understand Go's Error Handling"
  * Use multiple values in case conditions
  * How does the fallthrough statement work?
  * Solution: Parts of a Day
  * Recap: Switch Statement
  * How to continue a loop? (+BONUS: Debugging)
  * Create a multiplication table
  * How to loop over a slice?
  * For Range: Learn the easy way!

* **Projects: For Beginners** (Projetos para iniciantes)
  * Assista a todas as aulas.

* **Seções restantes**
  * A partir daqui, você pode assistir a todas as aulas restantes. São aulas de nível intermediário a avançado.

## É isso! Aproveite! 🤩

---

# BÔNUS: Por que você deveria aprender Go?

**Resumindo:** Go é fácil como Python e JavaScript e rápido como C/C++. Trabalhar com Go é mais agradável do que com C/C++. Você pode descer para o baixo nível ou ficar no alto nível.

## Para que o Go é usado?

Go é usado principalmente por empresas de internet: Google, Facebook, Twitter, Uber, Apple, Dropbox, Soundcloud, Medium, Mozilla Firefox, Github, Docker, Kubernetes e Heroku.

**Go é ideal para:** ferramentas de linha de comando multiplataforma, aplicações de rede distribuídas, tecnologias de nuvem como microsserviços e serverless, APIs web, motores de banco de dados, pipelines de processamento de big data, desenvolvimento embarcado e assim por diante.

**Go não é ideal para (mas dá para fazer):** aplicativos desktop, sistemas operacionais, drivers de kernel, desenvolvimento de jogos etc.

## Quem criou o Go?

O Go foi criado por algumas das pessoas mais influentes da indústria:

* Unix: Ken Thompson
* UTF-8, Plan 9: Rob Pike
* Hotspot JVM (Java Virtual Machine): Robert Griesemer

## QUANTO VOCÊ PODE GANHAR?

* [Salários de Go](https://www.payscale.com/research/US/Skill=Go_(Golang)_Programming_Language/Salary)

## [Do post "Eight years of Go"](https://blog.golang.org/8years):

> Hoje, **todas as empresas de nuvem têm componentes críticos da sua infraestrutura implementados em Go**, incluindo Google Cloud, AWS, Microsoft Azure, Heroku e muitas outras. Go é parte essencial de empresas de nuvem como Alibaba, Cloudflare e Dropbox. Go é parte crítica da infraestrutura aberta, incluindo Kubernetes, Cloud Foundry, Openshift, NATS, Docker, Istio, Etcd, Consul, Juju e muitos outros. Cada vez mais empresas escolhem Go para construir soluções de infraestrutura em nuvem.

## O que dá para fazer com Go?

* [Um driver de rede escrito em Go](https://www.net.in.tum.de/fileadmin/bibtex/publications/theses/2018-ixy-go.pdf) (_apenas 10% mais lento que um driver em C_)
* [Google gVisor](https://cloud.google.com/blog/products/gcp/open-sourcing-gvisor-a-sandboxed-container-runtime) (_kernel em espaço de usuário escrito em Go_)
* [Emulador de Nintendo multiplataforma](https://humpheh.github.io/goboy/)
* [Docker: sistema de contêineres](https://github.com/moby/moby)
* [Kubernetes: agendamento e gerenciamento de contêineres](https://github.com/kubernetes/kubernetes)
* Utilitário de deduplicação de imagens de VM
* Servidor de chat
* Coletor de beacons RUM
* Motor de banco de dados de séries temporais, com cliente, ferramentas de linha de comando etc.
* Biblioteca de map-reduce
* Proxy reverso em cluster para otimização de front-end, com reescrita de conteúdo em tempo real, redimensionamento de imagens, cache e execução de handlers de eventos em Lua (tudo multi-tenant)
* Nós de CDN com proxy reverso distribuídos geograficamente
* Daemon de gerenciamento de saúde com handlers de eventos e relatórios peer-to-peer
* Servidor DNS em Go puro
* Backend de API integrado ao MySQL
* Utilitário de captura/restauração de processos Linux
* Proxy reverso para esconder o servidor de assets dos clientes
* Conversor de HTML para PDF para geração de faturas
* Encurtador de URLs como tinyurl.com e goo.gl
* Serviço de mensagens SMS
* Gateway de pagamento com cartão de crédito
* Pacote de JSON Web Token
* Serviços de processamento de imagens em tempo real
* Render farm 3D/pipeline de produção de conteúdo (projeto bem grande)
* Implantação de contêineres lxc em produção
* Gerenciamento de testes automatizados

Referência: [este post no Reddit](https://www.reddit.com/r/golang/comments/5nac2b/what_have_you_used_go_for_in_your_professional/).

## PARA MAIS INFORMAÇÕES:

* [About Go: An Overview](https://blog.learngoprogramming.com/about-go-language-an-overview-f0bee143597c)
* [Why should you learn Go?](https://medium.com/@kevalpatel2106/why-should-you-learn-go-f607681fad65)
* [Emerging language of cloud Infrastructure](https://redmonk.com/dberkholz/2014/03/18/go-the-emerging-language-of-cloud-infrastructure/)
* [Empresas que usam Go](https://github.com/golang/go/wiki/GoUsers)
* [Eight years of Go](https://blog.golang.org/8years)
* [Twitter: Handling Five Billion Session in a Day with Go](https://blog.twitter.com/engineering/en_us/a/2015/handling-five-billion-sessions-a-day-in-real-time.html)
* [A C++ developer looks at Go](https://www.murrayc.com/permalink/2017/06/26/a-c-developer-looks-at-go-the-programming-language-part-1-simple-features/)

<div style="page-break-after: always;"></div>

> Mais tutoriais: [https://blog.learngoprogramming.com](https://blog.learngoprogramming.com)
>
> Copyright © 2019 Inanc Gumus
>
> Curso Aprenda Programação em Go
>
> [Clique aqui para ler a licença.](https://creativecommons.org/licenses/by-nc-sa/4.0/)
