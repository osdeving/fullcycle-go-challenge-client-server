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
CREATE TABLE IF NOT EXISTS cotacoes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    bid TEXT NOT NULL,
    ask TEXT NOT NULL,
    timestamp INTEGER NOT NULL,
    create_date TEXT NOT NULL
);
```

### 📌 **Descrição das Colunas**
| Coluna       | Tipo     | Descrição |
|-------------|---------|-----------|
| `id`        | INTEGER | Identificador único (auto incremento) |
| `bid`       | TEXT    | Valor de compra do dólar |
| `ask`       | TEXT    | Valor de venda do dólar |
| `timestamp` | INTEGER | Timestamp UNIX do momento da cotação |
| `create_date` | TEXT  | Data da cotação no formato `YYYY-MM-DD HH:MM:SS` |

## 🔄 Operações Principais

### ✅ **Inserir uma nova cotação**
```sql
INSERT INTO cotacoes (bid, ask, timestamp, create_date)
VALUES (?, ?, ?, ?);
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
- **As datas são armazenadas no formato `TEXT`** para compatibilidade com operações SQL.
- **O timestamp é armazenado como `INTEGER`** para facilitar comparações.
- **O banco de dados é leve** e armazena apenas cotações recentes para evitar sobrecarga.

🚀 **Agora seu banco de dados está bem estruturado e documentado!**

