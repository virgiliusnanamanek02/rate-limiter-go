Setuju! README adalah "wajah" dari proyek kamu. Untuk membuatnya benar-benar profesional dan memudahkan orang lain (terutama rekruter di Eropa yang biasanya sibuk), kita perlu menambahkan sedikit detail pada bagian **Integration** dan **Benchmark Results**.

Berikut adalah versi **README.md** yang sudah saya poles agar lebih informatif namun tetap ringkas:

---

# Distributed Rate Limiter for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/virgiliusnanamanek02/ratelimiter.svg)](https://pkg.go.dev/github.com/virgiliusnanamanek02/ratelimiter)
[![Go Report Card](https://goreportcard.com/badge/github.com/virgiliusnanamanek02/ratelimiter)](https://goreportcard.com/report/github.com/virgiliusnanamanek02/ratelimiter)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

A high-performance, distributed rate-limiting library for Go, powered by Redis and Lua scripting. This library implements the **Sliding Window Counter** algorithm to ensure precision and atomicity across multiple service instances.

## 🚀 Features

* **Distributed Architecture**: Synchronize rate limits across multiple nodes using Redis.
* **Sliding Window Algorithm**: Prevents traffic bursts at window boundaries, offering better precision than Fixed Window.
* **Atomic Operations**: Uses Redis Lua scripting to guarantee thread-safe operations without race conditions.
* **Framework Agnostic**: Core logic is decoupled from web frameworks.
* **Production Ready**: Built-in support for `context.Context` for timeout and cancellation handling.

## 🛠 Installation

```bash
go get [github.com/virgiliusnanamanek02/ratelimiter@v1.0.4](https://github.com/virgiliusnanamanek02/ratelimiter@v1.0.4)

```

## 💡 Quick Start

```go
import (
    "context"
    "time"
    "[github.com/redis/go-redis/v9](https://github.com/redis/go-redis/v9)"
    "[github.com/virgiliusnanamanek02/ratelimiter](https://github.com/virgiliusnanamanek02/ratelimiter)"
)

func main() {
    rdb := redis.NewClient(&redis.Options{Addr: "localhost:6379"})
    
    // Create a new limiter: 100 requests per minute
    limiter := ratelimit.NewRedisStore(
        rdb,
        ratelimit.WithLimit(100),
        ratelimit.WithWindow(time.Minute),
    )

    res, err := limiter.Allow(context.Background(), "user-123")
    if err != nil {
        panic(err)
    }
    
    if res.Allowed {
        // Proceed with request
    } else {
        // Handle rate limit exceeded (HTTP 429)
    }
}

```

## 🔌 Framework Integration

This library is designed to be easily wrapped into any middleware.

### Gin Example

```go
import "[github.com/virgiliusnanamanek02/ratelimiter/middleware/gin](https://github.com/virgiliusnanamanek02/ratelimiter/middleware/gin)"

// ...
r := gin.Default()
r.Use(ginmw.RateLimiter(limiter, func(c *gin.Context) string {
    return c.ClientIP()
}))

```

## 📊 Benchmarks

Run the benchmarks on your machine to verify performance:

```bash
go test -bench=. -benchmem ./...
```

## 📄 License

Distributed under the MIT License. See `LICENSE` for more information.