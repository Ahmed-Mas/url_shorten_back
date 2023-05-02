package serializer

import (
	"encoding/json"

	"github.com/Ahmed-Mas/url_shorten_back/pkg/shorten"
)

type Serializer struct{}

func (j *Serializer) Decode(input []byte) (*shorten.UrlData, error) {
	var urlData shorten.UrlData
	if err := json.Unmarshal(input, &urlData); err != nil {
		return nil, err
	}

	return &urlData, nil
}

func (j *Serializer) Encode(urlData *shorten.UrlData) ([]byte, error) {
	data, err := json.Marshal(urlData)
	if err != nil {
		return nil, err
	}
	return data, nil
}
