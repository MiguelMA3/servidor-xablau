# Servidor Xablau

## Descrição do Projeto

O ***Servidor Xablau*** é um programa escrito em Go que simula um servidor Amazon S3 simplificado. Ele permite que os usuários realizem operações básicas de armazenamento de arquivos, como upload, download e exclusão, por meio de requisições HTTP.

Este projeto foi desenvolvido como parte de um trabalho acadêmico para a disciplina de Sistemas Distribuídos. O objetivo é entender os conceitos de comunicação de rede, manipulação de arquivos e programação de servidores.

## Funcionamento do Servidor

O servidor utiliza a linguagem de programação Go e responde às seguintes requisições HTTP:

- GET: Para obter um arquivo do servidor.
- POST: Para enviar um arquivo para o servidor.
- DELETE: Para excluir um arquivo do servidor.

O servidor armazena os arquivos em uma pasta local chamada `storage/` no mesmo diretório do código-fonte.

## Teste

Para testar o servidor basta realizar uma requisição GET com a url seguida do nome do arquivo.

**Linux**
`curl http://localhost:8080/aquela.txt`

**Windows**
`Invoke-WebRequest -Uri "http://localhost:8080/aquela.txt"`

## Como Utilizar

1. **Configuração do Ambiente**:
   - Certifique-se de ter o Go instalado em seu sistema.
   - Clone ou baixe este repositório em sua máquina.

2. **Execução do Servidor**:
   - Abra um terminal e navegue até o diretório do projeto.
   - Execute o comando `go run server.go` ou `go run .` para iniciar o servidor.

3. **Testando as Requisições**:
   - Use um cliente HTTP, como o Postman, para enviar requisições ao servidor.
   - Envie requisições GET, POST e DELETE para interagir com os arquivos no servidor.

## Estrutura do Projeto

- `server.go`: Contém o código-fonte do servidor em Go.
- `storage/`: Pasta local para armazenamento dos arquivos.
- `README.md`: Este arquivo com informações detalhadas sobre o projeto.

## Autores

Este projeto foi desenvolvido por Miguel Maletzke e Victor Rey como parte da disciplina de Sistemas Distribuídos.