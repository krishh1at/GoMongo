package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	UserId string
	jwt.StandardClaims
}

func GetJWT(userId string) (string, error) {
	key := []byte(os.Getenv("SECRET_KEY"))

	claims := JWTClaims{
		userId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(key)

	if err != nil {
		log.Panicln("Error while generating JWT token: ", err)
		return "", err
	}

	return signedToken, nil
}

func VerifyJWT(token string) (bool, error) {
	key := []byte(os.Getenv("SECRET_KEY"))

	signedToken, err := jwt.ParseWithClaims(token, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("Someone tried to hijack JWT token")
		}

		return []byte(key), nil
	})

	if err != nil {
		return false, fmt.Errorf("Error while parsing JWT in validateJWT: %w", err)
	}

	return signedToken.Valid, nil
}
