package utils

import (
	"golang-jwt/config"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func EncryptToken(email, password string, expired int) (string, bool) {
	secretJWT := config.SecretJwt
	client := jwt.MapClaims {}
	client["email"] = email
	client["password"] = password
	client["exp"] = time.Now().Add(time.Second * time.Duration(expired)).Unix()
	encrypt := jwt.NewWithClaims(jwt.SigningMethodHS256, client)
	token, err := encrypt.SignedString([]byte(secretJWT))
	if err != nil {
		return "", true
	}
	return token, false
}

func DecryptToken(tokenStr string) (jwt.MapClaims, bool) {
	secretJWT := config.SecretJwt
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretJWT, nil
	})

	if err != nil {
		return nil, true
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, false
	} else {
		return nil, true
	}
}