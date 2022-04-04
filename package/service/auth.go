package service

import (
	"crypto/sha1"
	"fmt"
	"time"

	serv "github.com/Shin0kari/go_max"
	"github.com/Shin0kari/go_max/package/repository"
	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "sogbnsjkdfgn5645"
	signingKey = "g45g45dFSDF1d1541fd4"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

// структура дял работы с бд
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

// логика создания пользователей
func (s *AuthService) CreateUser(user serv.User) (int, error) {
	// хешируем пароль, а потом передаём в реп
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	// get user from DB
	user, err := s.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

// функция хеширования паролей
func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
