package user

import (
	"math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func tokenGenerator() Token {
	const lenght = 16
	rand.Seed(time.Now().UTC().UnixNano())
	bytes := make([]rune, lenght)
	for i := 0; i < lenght; i++ {
		bytes[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return Token(bytes)
}
