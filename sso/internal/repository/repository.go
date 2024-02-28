package repository

import "github.com/jmoiron/sqlx"

type Repostory struct {
	AuthRepoI
}

type AuthRepoI interface {
	Register(Email, Password string) (int, error)
	GetId(Email string) (int, error)
}

func NewRepository(db *sqlx.DB) *Repostory {
	return &Repostory{AuthRepoI: NewAuthRepo(db)}
}
