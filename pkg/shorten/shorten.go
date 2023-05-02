package shorten

import (
	"math/rand"
)

const validChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

type Shorten struct {
	storage Storage
}

func NewShorten(s Storage) *Shorten {
	return &Shorten{
		storage: s,
	}
}

func (s *Shorten) generateUrl() string {
	newUrl := []byte{}
	for i := 0; i < 7; i += 1 {
		idx := rand.Intn(61)
		newUrl = append(newUrl, []byte(validChars)[idx])
	}
	return string(newUrl)
}

func (s *Shorten) Retrieve(input string) *UrlData {
	return s.storage.Retrieve(input)
}

func (s *Shorten) Store(input *UrlData) {
	input.ShortUrl = s.generateUrl()
	s.storage.Store(input)
}
