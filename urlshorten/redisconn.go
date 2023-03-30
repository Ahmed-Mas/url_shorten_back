package urlshorten

type FakeRedis struct {
	redisConn map[string]string
}

func (redis *FakeRedis) fakeConnect() {
	if redis.redisConn == nil {
		redis.redisConn = make(map[string]string)
	}
}

func (redis *FakeRedis) PublishURL(holder *UrlHolder) int {
	redis.fakeConnect()
	redis.redisConn[holder.ShortURL] = holder.LongURL
	return 200
}

func (redis *FakeRedis) RetrieveURL(holder *UrlHolder) string {
	redis.fakeConnect()
	return redis.redisConn[holder.ShortURL]
}
