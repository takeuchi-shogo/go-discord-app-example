package worker

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/takeuchi-shogo/go-worker-scheduler-app-example/backend/events"
)

const (
	queueName      = "test_queue"
	maxWorkers     = 100
	maxQueueLength = 10000
)

var (
	wg  sync.WaitGroup
	ctx = context.Background()
)

type Job struct {
	ID      string
	Payload string
}

type Worker struct {
	client *redis.Client
}

func createChannels() []string {
	return events.RegisterEvents()
}

func New() {
	w := Worker{
		client: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", "localhost", "6379"),
		}),
	}
	// ジョブキューの初期化
	jobQueue := make(chan Job, maxQueueLength)

	// ワーカープールの初期化
	for i := 0; i < maxWorkers; i++ {
		go w.worker(jobQueue)
	}

	channels := createChannels()

	// 各チャンネルに対してサブスクライバーを起動
	for _, channel := range channels {
		go w.subscriberWorker(channel, jobQueue)
	}

	// メインゴルーチンをブロック
	wg.Wait()
}

func (w Worker) subscriberWorker(channel string, jobQueue chan<- Job) {
	pubsub := w.client.Subscribe(ctx, channel)
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			log.Printf("Error receiving message from channel %s: %v", channel, err)
			continue
		}

		job := Job{
			ID:      fmt.Sprintf("%s_%d", channel, time.Now().UnixNano()),
			Payload: msg.Payload,
		}

		// ジョブをキューに追加
		select {
		case jobQueue <- job:
			// ジョブがキューに追加された
		default:
			// キューが満杯の場合、ジョブをRedisのリストに保存
			err := w.client.LPush(ctx, queueName, job).Err()
			if err != nil {
				log.Printf("Error pushing job to Redis: %v", err)
			}
			fmt.Println("キューが満杯です")
		}
	}
}

func (w Worker) worker(jobQueue <-chan Job) {
	for {
		select {
		case job := <-jobQueue:
			processJob(job)
		default:
			// ジョブキューが空の場合、Redisからジョブを取得
			result, err := w.client.RPop(ctx, queueName).Result()
			if err == redis.Nil {
				// キューが空の場合は少し待機
				time.Sleep(100 * time.Millisecond)
			} else if err != nil {
				log.Printf("Error popping job from Redis: %v", err)
			} else {
				var job Job
				// 結果をJob構造体にアンマーシャル
				// エラーハンドリングは省略
				job.Payload = result
				processJob(job)
			}
		}
	}
}

func processJob(job Job) {
	// ジョブ処理のロジックをここに実装
	fmt.Printf("Processing job: %s with payload: %s\n", job.ID, job.Payload)
	// 実際の処理をここに追加
}
