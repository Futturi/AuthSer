package service

import "github.com/Futturi/AuthSer/sso/internal/repository"

type Service struct {
	Auth
}

type Auth interface {
	Register(Email, Password string) (int, error)
	Login(Email, Password string) (string, error)
}

func NewService(repo *repository.Repostory) *Service {
	return &Service{Auth: NewAuthService(repo.AuthRepoI)}
}
