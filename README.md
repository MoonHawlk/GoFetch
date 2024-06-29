# GoFetch

**GoFetch** é uma aplicação CLI escrita em Go que exibe informações do sistema de forma elegante no terminal, semelhante ao Neofetch. Ele fornece detalhes sobre o sistema operacional, plataforma, versão do kernel, tempo de atividade, memória total e usada, CPU e GPU.

## Índice

- [Funcionalidades](#funcionalidades)
- [Pré-requisitos](#pré-requisitos)
- [Instalação](#instalação)
- [Uso](#uso)
- [Contribuição](#contribuição)
- [Licença](#licença)

## Funcionalidades

- Exibe informações do sistema operacional.
- Mostra a plataforma e versão do kernel.
- Exibe o tempo de atividade do sistema.
- Mostra a quantidade total e usada de memória.
- Exibe detalhes da CPU e GPU.

## Pré-requisitos

Antes de começar, certifique-se de ter o Go instalado em sua máquina. Você pode baixar e instalar o Go a partir do [site oficial](https://golang.org/dl/).

Além disso, este projeto utiliza as seguintes bibliotecas Go:

- [gopsutil](https://github.com/shirou/gopsutil)
- [color](https://github.com/fatih/color)
- [try](https://github.com/matryer/try)

## Instalação

1. Clone o repositório para sua máquina local:

    ```sh
    git clone https://github.com/username/repository-name.git
    cd repository-name
    ```

2. Instale as dependências do projeto:

    ```sh
    go get github.com/shirou/gopsutil/v3
    go get github.com/fatih/color
    go get github.com/matryer/try
    ```

3. Compile o código fonte:

    ```sh
    go build -o gofetch main.go
    ```

## Uso

Execute o comando compilado para exibir as informações do sistema no terminal:

```sh
./gofetch
```

O resultado será algo assim:
```
+---------------------------------------------------+
| OS: windows                                       |
| Platform: Microsoft Windows 11 Pro                |
| Kernel Version: 10.0.22631 Build 22631            |
| Uptime: 9 hours                                   |
| Total Memory: 32693 MB                            |
| Used Memory: 10702 MB                             |
| CPU: AMD Ryzen 5 5600X 6-Core Processor           |
| GPU: NVIDIA GeForce RTX 4060 Ti                   |
+---------------------------------------------------+
```
## Contribuição
Contribuições são bem-vindas! Para contribuir, siga os seguintes passos:

1. Fork o repositório.
2. Crie um branch para sua feature (git checkout -b feature/feature-name).
3. Commit suas mudanças (git commit -m 'Add some feature').
4. Push para o branch (git push origin feature/feature-name).
5. Abra um Pull Request.

## Licença
Este projeto está licenciado sob a Licença MIT - veja o arquivo LICENSE para mais detalhes.

Feito por Filipe Moreno (2024) 💾.
