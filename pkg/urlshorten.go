package pkg

import (
	"math/rand"
)

const validChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func ShortenURL() string {
	newUrl := []byte{}
	for i := 0; i < 7; i += 1 {
		idx := rand.Intn(61)
		newUrl = append(newUrl, []byte(validChars)[idx])
	}

	return string(newUrl)
}
