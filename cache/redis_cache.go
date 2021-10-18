package cache

import (
	"encoding/json"
	"kumparan/models"
	"time"

	"github.com/go-redis/redis"
)

type redisCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewRedisCache(host string, db int, exp time.Duration) ArticleCache {
	return &redisCache{
		host:    host,
		db:      db,
		expires: exp,
	}
}

func (cache *redisCache) getClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cache.host, //127.0.0.1:6379
		Password: "",
		DB:       cache.db,
	})
}

func (cache *redisCache) Set(key string, value *models.Articles) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}

	client.Set("NEWS"+key, json, cache.expires*time.Second)
}

func (cache *redisCache) Get(key string) *models.Articles {
	client := cache.getClient()

	val, err := client.Get("NEWS" + key).Result()
	if err != nil {
		return nil
	}

	articles := models.Articles{}
	err = json.Unmarshal([]byte(val), &articles)
	if err != nil {
		panic(err)
	}
	return &articles
}
