package urlshorten

import (
	"log"
	"math/rand"
	"net/url"
)

const validUrlChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func cleanURL(longURL string) (string, error) {
	log.Println("cleaning url")
	validatedURL, err := url.ParseRequestURI(longURL)
	if err != nil {
		validatedURL, err = url.ParseRequestURI("https://" + longURL)
		if err != nil {
			return "", err
		}
	}
	log.Println("url clean")
	return validatedURL.String(), nil
}

func shortenHelper(longUrl string) string {
	newUrl := []byte{}
	for i := 0; i < 10; i += 1 {
		idx := rand.Intn(61)
		newUrl = append(newUrl, []byte(validUrlChars)[idx])
	}

	return string(newUrl)
}

func ShortenURL(longURL string) (string, error) {
	log.Printf("shortening long url: %s\n", longURL)
	longURL, err := cleanURL(longURL)
	if err != nil {
		return "", err
	}

	shortURL := shortenHelper(longURL)
	log.Printf("generated short url: %s\n", shortURL)
	return shortURL, nil
}
