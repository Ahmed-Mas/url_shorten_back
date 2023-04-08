package urlshorten

import (
	"math/rand"
)

const validUrlChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func ShortenURL() string {
	newUrl := []byte{}
	for i := 0; i < 10; i += 1 {
		idx := rand.Intn(61)
		newUrl = append(newUrl, []byte(validUrlChars)[idx])
	}

	return string(newUrl)
}
