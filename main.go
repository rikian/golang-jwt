package main

import (
	"golang-jwt/utils"
	"log"

	"github.com/google/uuid"
)

type User struct {
	id			 string
	Name     string
	Password string
}

func main() {
	var user User
	user.id = uuid.NewString()
	user.Name = "yahoo@gmail.com"
	user.Password = "r4h4514..."

	// encrypt token, expired in second
	token, statusToken := utils.EncryptToken(user.Name, user.Password, 60)
	if statusToken {
		log.Panicln(statusToken)
		return
	}
	log.Println(token)

	// decrypt token
	decryptToken, err := utils.DecryptToken(token)
	if err {
		log.Println("token expaired")
		return
	}
	log.Println(decryptToken)
}

