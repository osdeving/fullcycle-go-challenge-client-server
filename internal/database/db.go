package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/osdeving/fullcycle-go-challenge-client-server/pkg/models"
)

func ConnectDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./cotations.db")
	if err != nil {
		return nil, err
	}

	// Cria a tabela de usuários
	query := `
	CREATE TABLE IF NOT EXISTS cotations (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bid VARCHAR(20) NOT NULL,
		timestamp Timestamp NOT NULL,
		create_date Timestamp NOT NULL
	);
	`

	_, err = db.Exec(query)
	if err != nil {
		return nil, fmt.Errorf("❌ Erro ao criar tabela: %v", err)
	}

	return db, nil
}
func InsertCotation(ctx context.Context, db *sql.DB, moeda models.Moeda) error {

	query := `
		INSERT INTO cotations (bid, timestamp, create_date)
		VALUES ($1, $2, $3)
	`
	_, err := db.ExecContext(ctx, query, moeda.Bid, moeda.Timestamp, moeda.CreateDate)

	if err != nil {
		return fmt.Errorf("❌ Erro ao inserir cotação: %v", err)
	}

	log.Printf("✅ Cotação inserida com sucesso: %v", moeda)
	return nil
}
