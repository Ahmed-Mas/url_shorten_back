package redis

import "github.com/Ahmed-Mas/url_shorten_back/pkg/shorten"

// heres where ill put redis connection and the redis struct

type Redis struct {
	// Client: *redis.Client
}

func NewConn() *Redis {
	r := &Redis{}
	// client := redis.Client.Connect(redisUrl)
	// r.Client = client
	return r
}

func (r *Redis) Store(input *shorten.UrlData) {

}

func (r *Redis) Retrieve(input string) *shorten.UrlData {
	return &shorten.UrlData{}
}
