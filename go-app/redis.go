package main

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	RDB *redis.Client
	Ctx = context.Background()
)

func InitRedisSentinel() error {
	RDB = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName: "mymaster",

		// 🔥 Sentinel addresses (HOST MODE)
		SentinelAddrs: []string{
			"localhost:26379",
			"localhost:26380",
			"localhost:26381",
		},

		// 🔐 Redis AUTH (hardcoded)
		Password: "redis123",

		// Timeouts
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,

		// Pool tuning
		PoolSize:     20,
		MinIdleConns: 5,
		MaxRetries:   3,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := RDB.Ping(ctx).Err(); err != nil {
		return err
	}

	log.Println("[REDIS] connected via Sentinel")
	return nil
}

func CloseRedis() {
	if RDB != nil {
		_ = RDB.Close()
	}
}
