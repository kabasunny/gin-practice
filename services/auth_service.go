package services

import (
	"fmt"
	"gin-practice/models"
	"gin-practice/repositories"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Signup(email string, password string) error
	Login(email string, password string) (*string, error)
	GetUserFromToken(tokenString string) (*models.User, error)
}

type AuthService struct {
	repository repositories.IAuthRepository
}

func NewAuthService(repository repositories.IAuthRepository) IAuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) Signup(email string, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //文字列はイミュータブルで、バイトスライスに変換することで、ミュータブル化。また、ハッシュ関数はバイナリデータを処理するため、パスワードをバイトスライスに変換することで、処理できる
	//コストパラメータは、デフォルトだと10(範囲4～31)
	if err != nil {
		return err
	}
	user := models.User{
		Email:    email,
		Password: string(hashedPassword),
	}
	return s.repository.CreateUser(user)
}

func (s *AuthService) Login(email string, password string) (*string, error) {
	foundUser, err := s.repository.FindUser(email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	token, err := CreateToken(foundUser.ID, foundUser.Email)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func CreateToken(userId uint, email string) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   userId,
		"email": email,
		"exp":   time.Now().Add(time.Hour).Unix(),
	}) // Unixタイムスタンプは、1970年1月1日のUTC午前0時0分0秒からの経過秒数を表すため、int64型では約292億年分の秒数をカバー
	fmt.Printf("Token: %v\n", token) // 中間状態をコンソールで確認
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, err
	}
	return &tokenString, nil
}

func (s *AuthService) GetUserFromToken(tokenString string) (*models.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	var user *models.User
	// claimsにJWTトークンのクレームはinterface{}型で格納される
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			return nil, jwt.ErrTokenExpired
		}
		user, err = s.repository.FindUser(claims["email"].(string))
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}
