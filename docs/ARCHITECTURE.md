# 📌 Arquitetura do Projeto

Este documento descreve a arquitetura do projeto, incluindo sua estrutura de diretórios, fluxos de dados e principais tecnologias utilizadas.

## 📂 Estrutura de Diretórios

```
.
├── README.md            # 📌 Documentação principal do projeto
├── docs/                # 📂 Documentação detalhada
│   ├── API.md           # 📌 Documentação da API
│   ├── DATABASE.md      # 📌 Estrutura do banco de dados
│   ├── DEPLOY.md        # 📌 Guia de deploy
│   ├── ARCHITECTURE.md  # 📌 Arquitetura do sistema
│   └── assets/          # 📂 Diagramas e imagens
├── cmd/                 # 📂 Entrypoints do projeto (client/server)
│   ├── client/main.go   # 📌 Código principal do cliente
│   └── server/main.go   # 📌 Código principal do servidor
├── internal/            # 📂 Código interno do projeto
│   ├── auth/            # 📌 Autenticação (se necessário)
│   ├── client/          # 📌 Lógica do cliente
│   ├── config/          # 📌 Configurações do projeto
│   ├── database/        # 📌 Conexão com o banco de dados
│   └── server/          # 📌 Lógica do servidor
├── pkg/                 # 📂 Pacotes reutilizáveis
│   ├── api/             # 📌 Definições de API
│   ├── models/          # 📌 Estruturas de dados
│   └── utils/           # 📌 Utilitários gerais
├── deploy/              # 📂 Configurações de deployment
├── scripts/             # 📂 Scripts auxiliares
└── web/                 # 📂 Interface web (se aplicável)
```

## 🛠️ Tecnologias Utilizadas
- **Linguagem:** Go (Golang)
- **Banco de Dados:** SQLite
- **Web Server:** `net/http`
- **Autenticação:** (Se necessário, pode ser JWT ou OAuth)
- **Contexto:** `context` para controle de timeout

## 🔄 Fluxo de Dados
- 1️⃣ **Cliente (`client.go`)** faz uma requisição `GET /cotacao` para o servidor.
- 2️⃣ **Servidor (`server.go`)** busca a cotação na API `https://economia.awesomeapi.com.br/json/last/USD-BRL` (timeout de `200ms`).
- 3️⃣ **Servidor armazena no banco** de dados SQLite (timeout de `10ms`).
- 4️⃣ **Servidor responde ao cliente** com a cotação (`bid`) em JSON.
- 5️⃣ **Cliente salva a cotação** no arquivo `cotacoes.txt`.

## 📌 Diagramas
### 📌 Diagrama de Sequência

O diagrama abaixo representa o fluxo de comunicação entre os componentes do sistema.

![Diagrama de Sequência](assets/sequence.png)

<details>
  <summary>⚠ Caso a imagem não carregue, clique aqui para ver o código-fonte do diagrama ou renderizar via servidor externo.</summary>

```plantuml
@startuml
!define plantuml.server https://www.plantuml.com/plantuml/png/
actor Cliente
participant "Servidor (Go Web Server)" as Servidor
participant "API Externa (AwesomeAPI)" as API
database "Banco de Dados (SQLite)" as DB

Cliente -> Servidor: GET /cotacao
activate Servidor

Servidor -> API: Solicita cotação do dólar (timeout: 200ms)
activate API
API -> Servidor: Retorna JSON com bid
deactivate API

Servidor -> DB: Salva cotação (timeout: 10ms)
activate DB
DB -> Servidor: Confirma inserção
deactivate DB

Servidor -> Cliente: Retorna JSON (bid) (timeout: 300ms)
deactivate Servidor

Cliente -> Cliente: Salva bid no arquivo cotacoes.txt

@enduml
```
</details>

### **Diagrama C4**

![Diagrama C4](docs/assets/archtecture.png)

<details>
  <summary>⚠ Caso a imagem não carregue, clique aqui para ver o código-fonte do diagrama ou renderizar via servidor externo.</summary>

```plantuml
@startuml
!define plantuml.server https://www.plantuml.com/plantuml/png
!include <C4/C4_Container>

title C4 Diagram - Arquitetura do Sistema

Person(client, "Cliente CLI", "Solicita cotação do dólar")
System_Boundary(server, "Servidor Go") {
    Container(serverApp, "Servidor", "Go Web Server", "Recebe requisição do cliente e consulta API externa")
    ContainerDb(db, "Banco de Dados", "SQLite", "Armazena as cotações do dólar")
}
System_Ext(api, "API Externa", "AwesomeAPI", "Fornece a cotação USD/BRL")

Rel(client, serverApp, "GET /cotacao", "HTTP")
Rel(serverApp, api, "Busca cotação do dólar", "HTTP, timeout 200ms")
Rel(serverApp, db, "Salva cotação", "SQL, timeout 10ms")
Rel(serverApp, client, "Retorna JSON com bid", "HTTP")
Rel(db, serverApp, "Armazena e consulta cotações", "SQL")

@enduml
```
</details>

## 🚀 Conclusão
Esta arquitetura garante:
✅ **Baixa latência** com timeouts configurados.
✅ **Persistência dos dados** via SQLite.
✅ **Facilidade de manutenção** com estrutura modularizada.

---

🚀 **Codado por Willams "osdeving" Sousa** em 22/02/2025 🚀
