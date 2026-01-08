package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

func connection() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Docker exposed port
		Password: "",               // no password set
		DB:       0,                // default DB
	})

	// 1️⃣ Test connection
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
	}

	fmt.Println("✅ Redis connected successfully")

	// 2️⃣ Write test
	err := rdb.Set(ctx, "healthcheck", "ok", 0).Err()
	if err != nil {
		log.Fatalf("❌ SET failed: %v", err)
	}

	// 3️⃣ Read test
	val, err := rdb.Get(ctx, "healthcheck").Result()
	if err != nil {
		log.Fatalf("❌ GET failed: %v", err)
	}

	fmt.Println("✅ Redis GET value:", val)
}
