package urlshorten

import "testing"

func TestRedisPub(t *testing.T) {
	redis := FakeRedis{}
	code := redis.PublishURL(
		&UrlHolder{
			LongURL:  "https://example.com",
			ShortURL: "https://shorten.llamaherd.net/123445",
		},
	)
	if code != 200 {
		t.Errorf("unexpected code: %d", code)
	}
}

func TestRedisGet(t *testing.T) {
	redis := FakeRedis{}
	holder := UrlHolder{
		LongURL:  "https://example.com",
		ShortURL: "https://shorten.llamaherd.net/123445",
	}

	code := redis.PublishURL(&holder)
	if code != 200 {
		t.Errorf("unexpected code: %d", code)
	}

	longUrl := redis.RetrieveURL(&holder)
	if longUrl != "https://example.com" {
		t.Errorf("unexpected url: %s", longUrl)
	}
}
