package service

import (
	"time"

	"github.com/Futturi/AuthSer/sso/internal/repository"
	"github.com/golang-jwt/jwt"
)

const (
	salt = "eojgnrwijnweijfweijfnweijfniwjenfiwnsiquw"
)

type AuthService struct {
	repo repository.AuthRepoI
}

func NewAuthService(repo repository.AuthRepoI) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) Register(Email, Password string) (int, error) {
	return a.repo.Register(Email, Password)
}

type Claim struct {
	id int
	jwt.StandardClaims
}

func (a *AuthService) Login(Email, Password string) (string, error) {
	id, err := a.repo.GetId(Email)
	if err != nil {
		return "", err
	}
	claim := Claim{id,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(10 * time.Hour).Unix(),
		}}
	jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return jwt.SignedString([]byte(salt))
}
