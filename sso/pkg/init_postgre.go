package pkg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	Username string
	Port     string
	Hostname string
	DB       string
	Password string
	SSLMode  string
}

func InitPostges(cfg Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Hostname, cfg.Port, cfg.Username, cfg.Password, cfg.DB, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	return db, err
}
