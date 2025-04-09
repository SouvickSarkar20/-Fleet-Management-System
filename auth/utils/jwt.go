package utils

import (
	"fleet/auth/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	//used to store the subject of the token and other metadata

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//creating a object of jwt with the specific algos and signing method
	return token.SignedString([]byte(config.JWT_SECRET))
	//signs the token with the secret key
}
