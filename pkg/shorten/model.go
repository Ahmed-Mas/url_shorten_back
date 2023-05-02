package shorten

type UrlData struct {
	ShortUrl string `json:"short_url"`
	LongUrl  string `json:"long_url"`
	Time     int32  `json:"time"`
}

type Serializer interface {
	Encode(input *UrlData) ([]byte, error)
	Decode(input []byte) (*UrlData, error)
}

type Storage interface {
	Store(input *UrlData)
	Retrieve(input string) *UrlData
}

type Shortener interface {
	Store(input *UrlData)
	Retrieve(input string) *UrlData
}
