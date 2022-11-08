package utils

import (
	"math/rand"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString(n int) string {
	rndString := make([]byte, n)
	for i := range rndString {
		rndString[i] = alphabet[rand.Int63() % int64(len(alphabet))]
	}
	return  string(rndString)
}
