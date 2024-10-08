package redis

import (
	"fmt"

	"github.com/redis/go-redis/v9"
)

func New(c Config) *redis.Client {
	r := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", c.Host, c.Port),
	})

	return r
}
