package main

import (
	"context"
	"log"
	"time"
)

func StartWorker(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("[WORKER] stopping")
			return

		case <-ticker.C:
			err := RDB.Set(ctx, "heartbeat", time.Now().String(), 10*time.Second).Err()
			if err != nil {
				log.Println("[REDIS ERROR]", err)
				continue
			}

			val, _ := RDB.Get(ctx, "heartbeat").Result()
			log.Println("[REDIS] heartbeat =", val)
		}
	}
}
