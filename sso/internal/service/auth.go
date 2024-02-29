package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/Futturi/AuthSer/sso/internal/repository"
	"github.com/golang-jwt/jwt"
)

const (
	salt  = "eojgnrwijnweijfweijfnweijfniwjenfiwnsiquw"
	salt2 = "orjgnoqemvkwbm;lrmkpoejafkmgt;drlhmoherjso;wsozdnbhkjesr"
)

type AuthService struct {
	repo repository.AuthRepoI
}

func NewAuthService(repo repository.AuthRepoI) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) Register(Email, Password string) (int, error) {
	return a.repo.Register(Email, hashPass(Password))
}

func hashPass(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%v", hash.Sum([]byte(salt2)))
}

type Claim struct {
	Id int
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

func (a *AuthService) CheckIdentity(Header string) (int, error) {
	token, err := jwt.ParseWithClaims(Header, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return 0, errors.New("invalid signing method")
		}
		return []byte(salt), nil
	})
	if err != nil {
		return 0, err
	}
	Claim, ok := token.Claims.(*Claim)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return Claim.Id, nil
}
