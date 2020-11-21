package main

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	opt, err := redis.ParseURL("redis://localhost:6379/0")
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)

	pubsub := rdb.Subscribe(ctx, "examplechannel")

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			log.Println("an error occurred")
		}

		log.Println("received message: " + msg.Payload)
	}
}
