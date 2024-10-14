package runner

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/takeuchi-shogo/go-worker-scheduler-app-example/backend/events"
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
	router.POST("/greeting/:param", func(ctx *gin.Context) {
		param := ctx.Param("param")
		var channel string
		var message string
		switch param {
		case "1":
			channel = events.HelloWorld
			message = "こんにちわ世界"
		case "2":
			channel = events.GoodEvening
			message = "こんにちわ"
		case "3":
			channel = events.GoodAfternoon
			message = "こんばんわ"
		}
		result, err := rdb.Publish(ctx, channel, message).Result()
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
