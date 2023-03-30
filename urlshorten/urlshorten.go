package urlshorten

import (
	"math/rand"
	"net/url"
)

const validUrlChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const redirectURL = "https://shorten.llamaherd.net/v1/short/"

type UrlHolder struct {
	LongURL  string
	ShortURL string
}

func (holder *UrlHolder) cleanURL() error {
	validatedURL, err := url.ParseRequestURI(holder.LongURL)
	if err != nil {
		validatedURL, err = url.ParseRequestURI("https://" + holder.LongURL)
		if err != nil {
			return err
		}
	}

	holder.LongURL = validatedURL.String()
	return nil
}

func (holder *UrlHolder) shortenHelper() {
	newUrl := []byte{}
	for i := 0; i < 20; i += 1 {
		idx := rand.Intn(61)
		newUrl = append(newUrl, []byte(validUrlChars)[idx])
	}

	holder.ShortURL = redirectURL + string(newUrl)
}

func (holder *UrlHolder) ShortenURL() error {
	err := holder.cleanURL()
	if err != nil {
		return err
	}

	holder.shortenHelper()
	return nil
}
