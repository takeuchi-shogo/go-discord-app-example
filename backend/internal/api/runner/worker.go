package runner

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/redis/go-redis/v9"
	"github.com/takeuchi-shogo/go-worker-scheduler-app-example/backend/pkg/worker"
)

type Worker struct {
	redisClient *redis.Client
}

var ctx = context.Background()

func New() *Worker {
	return &Worker{
		redisClient: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", "localhost", "6379"),
		}),
	}
}

func (w *Worker) Start(ctx context.Context) {
	worker.New()
	pubsub := w.redisClient.Subscribe(ctx, "test-tasks")
	defer pubsub.Close()

	log.Println("ワーカーが開始され、タスクをリッスンしています")

	for {
		select {
		case <-ctx.Done():
			log.Println("ワーカーが停止しました")
			return
		default:
			msg, err := pubsub.ReceiveMessage(ctx)
			if err != nil {
				log.Printf("メッセージ受信エラー: %v", err)
				continue
			}
			log.Printf("受信したメッセージ: %s", msg.Payload)
			w.processTask(msg.Payload)
		}
	}
}

func (w *Worker) processTask(taskJSON string) {
	log.Println("JSON", taskJSON)
}

func NewWorker() {
	w := New()
	ctx, cancel := context.WithCancel(context.Background())

	go w.Start(ctx)

	// シグナルを待ち受けるチャネルを作成
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// シグナルを待機
	<-sigCh

	// キャンセル関数を呼び出してワーカーを停止
	cancel()
	log.Println("プログラムが終了しました")
}
