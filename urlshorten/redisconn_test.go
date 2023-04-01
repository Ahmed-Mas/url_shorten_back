package urlshorten

import (
	"testing"
)

func TestRedisPub(t *testing.T) {
	redis := RedisConn{}
	code := redis.PublishURL("https://shorten.llamaherd.net/123445", "https://example.com")
	if code != 200 {
		t.Errorf("unexpected code: %d", code)
	}
}

func TestRedisGet(t *testing.T) {
	redis := RedisConn{}

	code := redis.PublishURL("https://shorten.llamaherd.net/123445", "https://example.com")
	if code != 200 {
		t.Errorf("unexpected code: %d", code)
	}

	longUrl := redis.RetrieveURL("https://shorten.llamaherd.net/123445")
	if longUrl != "https://example.com" {
		t.Errorf("unexpected url: %s", longUrl)
	}
}
