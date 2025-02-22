# 📌 API - Cotacao do Dólar 💰

Este documento descreve os endpoints disponíveis no servidor e como utilizá-los.

## 🌍 Base URL
A API está disponível no seguinte endpoint:

```
http://localhost:8080
```

---

## 🔹 **1. Buscar Cotacao do Dólar**
### 📌 **GET /cotacao**
Obtém a cotação atual do dólar.

#### ✅ **Exemplo de Requisição**
```sh
curl -s "http://localhost:8080/cotacao" | jq
```

#### ✅ **Parâmetros**
| Parâmetro | Tipo   | Descrição |
|-----------|--------|-----------|
| `all`     | `bool` | (Opcional) Se `true`, retorna todas as informações. Se `false`, retorna apenas `bid`. |

#### ✅ **Respostas**
📌 **Se `all=true` (todas as informações):**
```json
{
  "code": "USD",
  "bid": "5.2100",
  "ask": "5.2200",
  "timestamp": "1708586152",
  "create_date": "2025-02-22 08:25:12"
}
```

📌 **Se `all=false` ou não for passado (apenas `bid`):**
```json
{
  "bid": "5.2100"
}
```

📌 **Se a API externa falhar (erro 500):**
```json
{
  "error": "Erro ao buscar cotacao"
}
```

📌 **Se o servidor demorar mais de 300ms (timeout):**
```json
{
  "error": "Timeout: O servidor demorou para responder"
}
```

---

## 🔹 **2. Como funciona o timeout?**
| Contexto | Timeout | O que acontece se exceder? |
|----------|---------|---------------------------|
| **Requisição à API externa** | `200ms` | Se a API não responder em 200ms, retorna `504 Gateway Timeout`. |
| **Resposta ao cliente** | `300ms` | Se a API demorar mais de 300ms, retorna `504 Gateway Timeout`. |
| **Inserção no banco de dados** | `10ms` | Se a inserção no SQLite demorar mais de 10ms, será cancelada e registrada como erro. |

---

## 📌 **Exemplo de uso com cURL**
1️⃣ **Buscar apenas o `bid`**
```sh
curl -s "http://localhost:8080/cotacao" | jq
```

2️⃣ **Buscar todas as informações**
```sh
curl -s "http://localhost:8080/cotacao?all=true" | jq
```

3️⃣ **Simular timeout (API externa demora mais de 200ms)**
```sh
curl -s "http://localhost:8080/cotacao"
# Espera erro 504 Gateway Timeout
```

---

## ✅ **Conclusão**
- 📌 O endpoint **`GET /cotacao`** retorna a cotação do dólar.
- 📌 Pode ser usado com `all=true` para obter mais detalhes.
- 📌 Tem **timeouts de segurança** para evitar travamentos.
- 📌 Todos os erros são tratados e retornam JSON.

---

🚀 **Codado por Willams "osdeving" Sousa** em 22/02/2025 🚀

