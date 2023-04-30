package pkg

import (
	"log"
)

type RedisConn struct {
	redisConn map[string]string
}

func (redis *RedisConn) fakeConnect() {
	if redis.redisConn == nil {
		log.Println("connecting to redis")
		redis.redisConn = make(map[string]string)
	}
}

func (redis *RedisConn) checkSize() {
	if len(redis.redisConn) >= 10 {
		redis.redisConn = make(map[string]string)
	}
}

func (redis *RedisConn) PublishURL(shortURL, longURL string) int {
	log.Printf("storing url in redis\n\tshort url: %s\n\tlong url: %s\n", shortURL, longURL)
	redis.fakeConnect()
	redis.checkSize()
	redis.redisConn[shortURL] = longURL
	return 200
}

func (redis *RedisConn) RetrieveURL(shortURL string) string {
	log.Printf("retrieving long url from redis\n\tshort url: %s\n", shortURL)
	redis.fakeConnect()
	return redis.redisConn[shortURL]
}
