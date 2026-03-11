package repository

import (
	"database/sql"
	"fmt"

	"subscriptions-service/internal/config"

	_ "github.com/lib/pq"
)

// NewDB возвращает *sql.DB для Postgres
func NewDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Проверка соединения
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
