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

		// ✅ HOST → Docker (ports published)
		SentinelAddrs: []string{
			"127.0.0.1:26379",
			"127.0.0.1:26380",
			"127.0.0.1:26381",
		},

		// ❌ REMOVE unless Redis AUTH is enabled
		// Password: "redis123",

		DialTimeout:  10 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,

		PoolSize:     20,
		MinIdleConns: 5,
		MaxRetries:   3,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
