package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/virgiliusnanamanek02/rate-limiter-go"
	ratelimitgin "github.com/virgiliusnanamanek02/rate-limiter-go/middleware/gin"
)

func main() {
	// 1. Inisialisasi Redis Client
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// 2. Inisialisasi Rate Limiter (e.g., 5 request per 10 detik)
	limiter := ratelimit.NewRedisStore(
		rdb,
		ratelimit.WithLimit(5),
		ratelimit.WithWindow(10*time.Second),
	)

	// 3. Setup Gin Engine
	r := gin.Default()

	// 4. Tentukan bagaimana cara kita mengidentifikasi user (misal lewat IP)
	keyFunc := func(c *gin.Context) string {
		return c.ClientIP()
	}

	// 5. Pasang Middleware ke Route tertentu atau Global
	r.Use(ratelimitgin.RateLimiter(limiter, keyFunc))

	// 6. Define Endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	log.Println("Server running on :8080")
	r.Run(":8080")
}