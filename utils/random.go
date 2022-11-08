package utils

import (
	"math/rand"
	"time"
)

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandomString_old(n int) string {
	rndString := make([]byte, n)
	for i := range rndString {
		rndString[i] = alphabet[rand.Int63() % int64(len(alphabet))]
	}
	return  string(rndString)
}


var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func RandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = alphabet[seededRand.Intn(len(alphabet))]
	}
	return string(b)
}