package runner

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func NewApiServer() {
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", "localhost", "6379"),
	})
	router := gin.Default()
	router.POST("/message", func(ctx *gin.Context) {
		result, err := rdb.Publish(ctx, "test-tasks", "HELLO!!!!!").Result()
		if err != nil {
			log.Println("エラー:", err)
			ctx.JSON(500, gin.H{"error": "メッセージの公開に失敗しました"})
			return
		}
		log.Printf("メッセージが公開されました。受信者数: %d", result)
		ctx.JSON(200, gin.H{"message": "メッセージが正常に公開されました", "receivers": result})
	})
	// NewWorker()
	router.Run(":8080")
}
