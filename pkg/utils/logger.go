package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/osdeving/fullcycle-go-challenge-client-server/pkg/models"
)

func SaveToFile(filename string, moeda models.Bid) {
	// 📌 Obtém timestamp ISO 8601
	timestamp := time.Now().Format(time.RFC3339)

	// 📌 Abre/cria o arquivo para escrita
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("❌ Erro ao abrir arquivo:", err)
		return
	}
	defer file.Close()

	// 📌 Escreve no arquivo com timestamp
	_, err = file.WriteString(fmt.Sprintf("[%s] Dólar: %s\n", timestamp, moeda.Bid))
	if err != nil {
		fmt.Println("❌ Erro ao escrever no arquivo:", err)
		return
	}

	fmt.Printf("✅ Cotação salva no arquivo %s com sucesso!\n", filename)
}
