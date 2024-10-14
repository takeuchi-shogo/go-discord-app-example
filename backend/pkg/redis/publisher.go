package redis

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
	"github.com/takeuchi-shogo/go-worker-scheduler-app-example/backend/pkg/database/entity"
)

type Publisher struct {
	client RedisClient
}

func NewPublisher(rc *redis.Client) Publisher {
	return Publisher{
		client: rc,
	}
}

func (p Publisher) PublishGreeting(ctx context.Context, channel string, message interface{}) error {
	return nil
}

func (p Publisher) PublishCreateUser(ctx context.Context, channel string, user entity.User) error {
	return p.publish(ctx, channel, user)
}

func (p Publisher) publish(ctx context.Context, channel string, event interface{}) error {
	var b bytes.Buffer

	if err := json.NewEncoder(&b).Encode(event); err != nil {
		return err
	}

	res := p.client.Publish(ctx, channel, b.Bytes())
	if err := res.Err(); err != nil {
		return err
	}

	return nil
}
