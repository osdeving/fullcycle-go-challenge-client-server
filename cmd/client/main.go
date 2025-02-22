package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/osdeving/fullcycle-go-challenge-client-server/internal/client"
	"github.com/osdeving/fullcycle-go-challenge-client-server/pkg/utils"
)

func main() {
	fmt.Println("Starting client...")

	// timeout (300ms)
	ctxCotacao, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	// Busca a cotação
	moeda, err := client.FetchCotacao(ctxCotacao)
	if err != nil {
		color.New(color.FgRed).Printf("\n❌ Erro: %v\n", err)
		return
	}

	// Exibe a cotação formatada
	utils.PrintCotacao(moeda)

	// timeout (100ms)
	ctxBid, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	// Busca apenas Bid
	bid, err := client.FetchBid(ctxBid)
	if err != nil {
		color.New(color.FgRed).Printf("\n\n\n❌ Erro: %v\n", err)
		return
	}

	fmt.Printf("\n\n\nBid: %v\n", bid.Bid)
	utils.SaveToFile("./cotacoes.txt", bid)
}
