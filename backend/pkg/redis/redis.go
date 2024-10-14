package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	redis.Cmdable
}

func New(c Config) RedisClient {
	r := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", c.Host, c.Port),
	})

	return r
}
