package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	log.Println("[APP] starting")

	if err := InitRedisSentinel(); err != nil {
		log.Fatal("[APP] redis init failed:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	go heartbeatWorker(ctx)

	// Graceful shutdown
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	log.Println("[APP] shutdown signal received")
	cancel()

	time.Sleep(2 * time.Second)
	CloseRedis()
	log.Println("[APP] shutdown complete")
}

func heartbeatWorker(ctx context.Context) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			log.Println("[WORKER] stopped")
			return

		case <-ticker.C:
			err := RDB.Set(
				ctx,
				"heartbeat",
				time.Now().Format(time.RFC3339),
				10*time.Second,
			).Err()

			if err != nil {
				log.Println("[REDIS ERROR]", err)
				continue
			}

			val, _ := RDB.Get(ctx, "heartbeat").Result()
			log.Println("[REDIS] heartbeat =", val)
		}
	}
}
