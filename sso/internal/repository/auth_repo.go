package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthRepo struct {
	db *sqlx.DB
}

func NewAuthRepo(db *sqlx.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (a *AuthRepo) Register(Email, Password string) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s(email, password) VALUES($1,$2) RETURNING id", "users")
	row := a.db.QueryRow(query, Email, Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (a *AuthRepo) GetId(Email string) (int, error) {
	var id int
	query := fmt.Sprintf("SELECT id FROM %s WHERE email = $1", "users")
	row := a.db.QueryRow(query, Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
