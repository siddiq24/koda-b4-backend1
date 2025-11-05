package utils

import (
	"log"

	"github.com/matthewhartstonge/argon2"
)

func Encoding(pass string) (string) {
	argon := argon2.DefaultConfig()

	encoded, err := argon.HashEncoded([]byte(pass))
	if err != nil {
		log.Println(err)
	}
	return string(encoded)
}

func Matching(usePass, dbPass string) bool {
	ok, err := argon2.VerifyEncoded([]byte(usePass), []byte(dbPass))
	if err != nil {
		log.Println(err)
		return false
	}
	if ok {
		return true
	}
	return  false
}