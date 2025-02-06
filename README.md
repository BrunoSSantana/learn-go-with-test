# Aprendendo Go com Testes

Este repositório contém meus exercícios e anotações enquanto estudo o livro [Aprenda Go com Testes](https://larien.gitbook.io/aprenda-go-com-testes), que ensina a linguagem Go através do desenvolvimento orientado a testes (TDD).

## Sobre o Livro

"Aprenda Go com Testes" é um material didático que explora a linguagem Go escrevendo testes. O Go é uma linguagem simples de aprender e possui suporte nativo para testes, tornando-o ideal para praticar TDD. O livro está disponível em:

- [Versão em inglês](https://quii.gitbook.io/learn-go-with-tests)
- [Versão em português](https://larien.gitbook.io/aprenda-go-com-testes)

## Estrutura do Repositório

O repositório segue a estrutura abaixo, onde cada diretório representa um conceito ou capítulo do livro, contendo o código-fonte e seus respectivos testes:

```
learn-go-with-test/
│── cmd/
│   ├── array_slices/
│   │   ├── sum.go
│   │   ├── sum_test.go
│   ├── dependency/
│   │   ├── dependency.go
│   │   ├── dependency_test.go
│   ├── hello-world/
│   │   ├── hello.go
│   │   ├── hello_test.go
│   ├── integers/
│   │   ├── adder.go
│   │   ├── adder_test.go
│   ├── iteration/
│   │   ├── iteration.go
│   │   ├── iteration_test.go
│   ├── maps/
│   │   ├── maps.go
│   │   ├── maps_test.go
│   ├── mocks/
│   │   ├── mocks.go
│   │   ├── mocks_test.go
│   ├── pointers/
│   │   ├── pointers.go
│   │   ├── pointers_test.go
│   ├── structs/
│   │   ├── structs.go
│   │   ├── structs_test.go
│── go.mod
│── main.go
```

Cada diretório contém:

- **`<conceito>.go`**: Implementação do conceito abordado no capítulo.
- **`<conceito>_test.go`**: Testes unitários escritos para validar o comportamento do código.

> O arquivo `main.go` pode ser utilizado para executar alguns dos exemplos de forma independente.

Para rodar os testes de um módulo específico, navegue até o diretório correspondente e execute:

```bash
go test
```


## Como Executar os Exemplos

1. **Clonar o repositório**:

   ```bash
   git clone https://github.com/seu-usuario/seu-repositorio.git
   cd seu-repositorio
   ```

2. **Navegar até o capítulo desejado**:

   ```bash
   cd 02-ola-mundo
   ```

3. **Executar o código**:

   ```bash
   go run main.go
   ```

4. **Executar os testes**:

   ```bash
   go test
   ```

> **Pré-requisitos**: Certifique-se de ter o [Go](https://golang.org/dl/) instalado em sua máquina.

## Contribuições

Este repositório é um registro do meu aprendizado pessoal, mas fico feliz em receber sugestões ou correções. Sinta-se à vontade para abrir uma issue ou enviar um pull request.

## Licença

Este projeto está licenciado sob a Licença MIT. Consulte o arquivo [LICENSE](./LICENSE) para mais detalhes.
