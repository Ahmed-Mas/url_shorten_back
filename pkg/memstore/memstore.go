package memstore

import (
	"log"

	"github.com/Ahmed-Mas/url_shorten_back/pkg/shorten"
)

type MemStore struct {
	store map[string]string
}

func NewMemStore() *MemStore {
	rediConn := &MemStore{}
	rediConn.initialize()
	return rediConn
}

func (m *MemStore) initialize() {
	if m.store == nil {
		log.Println("initializing empty map")
		m.store = make(map[string]string)
	}
}

func (m *MemStore) checkSize() {
	if len(m.store) >= 10 {
		m.store = make(map[string]string)
	}
}

func (m *MemStore) Store(input *shorten.UrlData) {
	log.Printf("storing url in mem\n\tshort url: %s\n\tlong url: %s\n", input.ShortUrl, input.LongUrl)
	m.checkSize()
	m.store[input.ShortUrl] = input.LongUrl
}

func (m *MemStore) Retrieve(input string) *shorten.UrlData {
	log.Printf("retrieving long url from mem\n\tshort url: %s\n", input)
	data := &shorten.UrlData{
		ShortUrl: input,
		LongUrl:  m.store[input],
	}
	return data
}
