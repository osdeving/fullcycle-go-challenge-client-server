package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/osdeving/fullcycle-go-challenge-client-server/internal/database"
	"github.com/osdeving/fullcycle-go-challenge-client-server/pkg/models"
)

var db *sql.DB

func Start() {
	fmt.Println("Server is running...")

	var err error
	db, err = database.ConnectDB()
	if err != nil {
		log.Fatalf("❌ Erro ao conectar ao banco de dados: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/cotacao", cotationHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("🚀 Servidor iniciado na porta 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("❌ Erro ao iniciar o servidor: %v", err)
	}
}

func cotationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("🔍 Buscando cotação...")

	// contexto (Timeout: 200ms)
	ctxAPI, cancelAPI := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancelAPI()

	req, err := http.NewRequestWithContext(ctxAPI, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		http.Error(w, "Erro ao criar requisição", http.StatusInternalServerError)
		log.Println("❌ Erro na criação da requisição:", err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		if ctxAPI.Err() == context.DeadlineExceeded {
			log.Println("⏳ ❌ Timeout: API demorou mais de 200ms para responder")
			http.Error(w, "❌ Timeout: API demorou para responder", http.StatusGatewayTimeout)
		} else {
			log.Println("❌ Erro ao chamar a API:", err)
			http.Error(w, "Erro ao chamar a API", http.StatusInternalServerError)
		}
		return
	}
	defer res.Body.Close()

	var cotacao models.Cotacao
	if err := json.NewDecoder(res.Body).Decode(&cotacao); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusInternalServerError)
		log.Println("❌ Erro ao decodificar JSON:", err)
		return
	}

	// validar query param "all=true"
	allParam := r.URL.Query().Get("all")
	var responseData any

	if allParam == "true" {
		responseData = cotacao.Moeda
		log.Println("📊 Enviando cotação completa")
	} else if allParam == "" || allParam == "false" {
		// retorna apenas o bid
		responseData = struct {
			Bid string `json:"bid"`
		}{Bid: cotacao.Moeda.Bid}
		log.Println("💰 Enviando apenas o valor do Bid")
	} else {
		http.Error(w, "Parâmetro inválido para 'all'", http.StatusBadRequest)
		log.Println("❌ Parâmetro inválido para 'all':", allParam)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		log.Println("❌ Erro ao enviar resposta para o cliente:", err)
	} else {
		log.Println("✅ Resposta enviada ao cliente com sucesso!")
	}

	// contexto (Timeout: 10ms)
	if err := insertCotationWithTimeout(cotacao.Moeda); err != nil {
		log.Println("❌ Falha ao salvar no banco:", err)
	}
}

// insere a cotação no banco
func insertCotationWithTimeout(cotacao models.Moeda) error {
	ctxDB, cancelDB := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancelDB()

	err := database.InsertCotation(ctxDB, db, cotacao)
	if err != nil {
		if ctxDB.Err() == context.DeadlineExceeded {
			log.Println("⏳ ❌ Timeout: Inserção no banco demorou mais de 10ms")
		} else {
			log.Println("❌ Erro ao salvar no banco:", err)
		}
		return err
	}

	log.Println("✅ Cotação salva no banco com sucesso!")
	return nil
}
