package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	entity "github.com/techlateef/jwt-Auth-techies/entities"
)

type JWTService interface {
	GeneratedToken(Role, userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

// this contains the payload Register claims
type jwtCustomClaim struct {
	Role   string `json:"role"`
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// this are for the signature generation
type jwtService struct {
	secretKey string
	issuer    string
}

// NewJWTService method is
func NewJWTService() JWTService {
	return &jwtService{
		issuer:    "ydhnwb",
		secretKey: getSecretKey(),
	}
}

// this func is resposible for getting thr secretKey from .env file
func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey != "" {
		secretKey = "ydhnwb"
	}
	return secretKey
}

// this generate the token using the the header whic is hs256, claims and secret key
func (j *jwtService) GeneratedToken(Role, userID string) string {
	var user entity.User
	claims := &jwtCustomClaim{

		user.Role,
		string(rune(user.ID)),
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKey), nil
	})
}
