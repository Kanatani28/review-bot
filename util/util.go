package util

import (
	"log"
	"math/rand"

	"github.com/pkg/errors"
)

func ErrLog(err error) {
	log.Printf("%+v\n", errors.WithStack(err))
}

func NewSecret(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func IndexOf(array []string, target string) int {
	for i, element := range array {
		if element == target {
			return i
		}
	}
	return -1
}
