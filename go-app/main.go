// package main

// import (
// 	"context"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"
// 	"time"
// )

// func main() {
// 	log.Println("[APP] starting")

// 	if err := InitRedisSentinel(); err != nil {
// 		log.Fatal("[APP] redis init failed:", err)
// 	}

// 	ctx, cancel := context.WithCancel(context.Background())
// 	go heartbeatWorker(ctx)

// 	// Graceful shutdown
// 	sig := make(chan os.Signal, 1)
// 	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
// 	<-sig

// 	log.Println("[APP] shutdown signal received")
// 	cancel()

// 	time.Sleep(2 * time.Second)
// 	CloseRedis()
// 	log.Println("[APP] shutdown complete")
// }

// func heartbeatWorker(ctx context.Context) {
// 	ticker := time.NewTicker(2 * time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			log.Println("[WORKER] stopped")
// 			return

// 		case <-ticker.C:
// 			err := RDB.Set(
// 				ctx,
// 				"heartbeat",
// 				time.Now().Format(time.RFC3339),
// 				10*time.Second,
// 			).Err()

// 			if err != nil {
// 				log.Println("[REDIS ERROR]", err)
// 				continue
// 			}

// 			val, _ := RDB.Get(ctx, "heartbeat").Result()
// 			log.Println("[REDIS] heartbeat =", val)
// 		}
// 	}
// }

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	connection()
}

func sentinal() {
	log.Println("[TEST] starting redis sentinel test")

	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName: "mymaster",

		// 🔥 Sentinel DNS names (Docker network)
		SentinelAddrs: []string{
			"sentinel-1:26379",
			"sentinel-2:26379",
			"sentinel-3:26379",
		},

		Password: "", // set if auth enabled

		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,

		PoolSize: 10,
	})

	ctx := context.Background()

	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatal("[FAIL] ping failed:", err)
	}

	log.Println("[OK] connected to Redis via Sentinel")

	// Write
	err := rdb.Set(ctx, "hello", "world", 0).Err()
	if err != nil {
		log.Fatal("set failed:", err)
	}

	// Read
	val, err := rdb.Get(ctx, "hello").Result()
	if err != nil {
		log.Fatal("get failed:", err)
	}

	fmt.Println("[OK] value from redis =", val)

	// Keep running to observe failover
	for i := 1; ; i++ {
		time.Sleep(2 * time.Second)
		err := rdb.Incr(ctx, "counter").Err()
		if err != nil {
			log.Println("[WARN] incr failed:", err)
			continue
		}
		fmt.Println("counter =", i)
	}
}
