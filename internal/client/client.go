package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/osdeving/fullcycle-go-challenge-client-server/pkg/models"
)

// Função genérica para requisição HTTP com timeout
func doConnection(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("❌ Erro ao criar requisição: %v", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("❌ Erro ao fazer requisição: %v", err)
	}

	// Verifica se o status code está dentro do intervalo válido (200-299)
	if res.StatusCode > 299 {
		res.Body.Close() // Fechar o corpo da resposta para evitar vazamento de memória
		return nil, fmt.Errorf("❌ Internal Server Error: API retornou %d", res.StatusCode)
	}

	return res, nil
}

// Função genérica para buscar e desserializar JSON
func fetchData[T any](ctx context.Context, url string) (T, error) {
	var result T

	res, err := doConnection(ctx, url)
	if err != nil {
		return result, err
	}
	defer res.Body.Close() // Fecha apenas se a conexão for bem-sucedida

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return result, fmt.Errorf("❌ Erro ao decodificar JSON: %v", err)
	}

	return result, nil
}

// Busca todas as cotações (USDBRL)
func FetchCotacao(ctx context.Context) (models.Moeda, error) {
	return fetchData[models.Moeda](ctx, "http://localhost:8080/cotacao?all=true")
}

// Busca apenas o valor do Bid
func FetchBid(ctx context.Context) (models.Bid, error) {
	return fetchData[models.Bid](ctx, "http://localhost:8080/cotacao")
}
