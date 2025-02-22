package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/osdeving/fullcycle-go-challenge-client-server/pkg/models"
)

// Formata e exibe a cotação no terminal de forma estilizada
func PrintCotacao(moeda models.Moeda) {
	clearScreen()

	// Estilos de cores
	headerStyle := color.New(color.FgWhite)
	infoStyle := color.New(color.FgGreen)
	warningStyle := color.New(color.FgYellow)
	errorStyle := color.New(color.FgRed)

	// Cabeçalho
	headerStyle.Println("\n💰 COTAÇÃO DO DÓLAR 💰")
	headerStyle.Println("────────────────────────────────────────")

	// Exibe informações formatadas
	infoStyle.Printf("📅 Data: %s\n", moeda.CreateDate)
	infoStyle.Printf("🔄 Última atualização: %s\n", time.Unix(atoi(moeda.Timestamp), 0).Format("15:04:05"))

	// Valores
	warningStyle.Printf("\n💵 Cotação atual\n")
	fmt.Printf("   🟢 Compra (Bid): R$ %s\n", moeda.Bid)
	fmt.Printf("   🔴 Venda (Ask): R$ %s\n", moeda.Ask)
	fmt.Printf("   📈 Alta do dia: R$ %s\n", moeda.High)
	fmt.Printf("   📉 Baixa do dia: R$ %s\n", moeda.Low)
	fmt.Printf("   🔁 Variação: %s%%\n", moeda.PctChange)

	// Rodapé
	headerStyle.Println("\n────────────────────────────────────────")
	errorStyle.Println("🚨 Valores podem variar dependendo da instituição financeira 🚨")
}

func atoi(s string) int64 {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return val
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
