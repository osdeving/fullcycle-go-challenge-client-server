# 📌 Estrutura do Banco de Dados

Este documento descreve a estrutura do banco de dados utilizado no projeto, incluindo tabelas, esquemas e operações principais.

## 🛠️ Banco de Dados Utilizado
- **Banco:** SQLite
- **Driver:** `github.com/mattn/go-sqlite3`
- **Arquivo de banco:** `cotations.db`

## 📂 Estrutura da Tabela

### 🔹 **Tabela: cotacoes**
Armazena as cotações do dólar obtidas da API externa.

```sql
CREATE TABLE IF NOT EXISTS cotations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid VARCHAR(20) NOT NULL,
		timestamp Timestamp NOT NULL,
		create_date Timestamp NOT NULL
	);
```

### 📌 **Descrição das Colunas**
| Coluna       | Tipo     | Descrição |
|-------------|--------------|--------------------------------------------------|
| `id`        | INTEGER      | Identificador único (auto incremento)            |
| `bid`       | VARCHAR(20)  | Valor de compra do dólar                         |
| `timestamp` | TIMESTAMP    | Timestamp UNIX do momento da cotação             |
| `create_date` | TIMESTAMP  | Data da cotação no formato `YYYY-MM-DD HH:MM:SS` |

## 🔄 Operações Principais

### ✅ **Inserir uma nova cotação**
```sql
INSERT INTO cotacoes (bid, timestamp, create_date)
VALUES (?, ?, ?);
```

### ✅ **Buscar as últimas cotações**
```sql
SELECT * FROM cotacoes ORDER BY timestamp DESC LIMIT 10;
```

### ✅ **Excluir cotações antigas (exemplo: manter últimos 1000 registros)**
```sql
DELETE FROM cotacoes WHERE id NOT IN (
    SELECT id FROM cotacoes ORDER BY timestamp DESC LIMIT 1000
);
```

## 📌 Considerações
- **O ID é gerado automaticamente** pelo SQLite (`AUTOINCREMENT`).
- **O banco de dados é o sqlite3** e armazena em um arquivo.

---

🚀 **Codado por Willams "osdeving" Sousa** em 22/02/2025 🚀

