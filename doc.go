// Package ratelimiter provides a distributed rate limiter for Go
// using Redis and the Sliding Window Counter algorithm.
//
// This package is designed for high-performance, concurrent systems
// and ensures atomic rate-limiting operations across multiple instances
// by leveraging Redis Lua scripting.
//
// Typical use cases include:
//   - API rate limiting
//   - Distributed systems throttling
//   - Per-user or per-IP request control
//
// See the README for integration examples and framework-specific middleware.
package ratelimit
