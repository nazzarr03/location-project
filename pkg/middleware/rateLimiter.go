package middleware

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	rateLimit  = 10
	ratePeriod = 60 * time.Second
	requests   = make(map[string][]time.Time)
	mu         sync.Mutex
)

func RateLimiter() fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		mu.Lock()

		timestamps, ok := requests[ip]
		if !ok {
			timestamps = []time.Time{}
		}

		cutoff := time.Now().Add(-ratePeriod)
		for i := 0; i < len(timestamps); i++ {
			if timestamps[i].Before(cutoff) {
				timestamps = append(timestamps[:i], timestamps[i+1:]...)
				i--
			}
		}

		if len(timestamps) >= rateLimit {
			mu.Unlock()
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"message": "Too many requests",
			})
		}

		timestamps = append(timestamps, time.Now())
		requests[ip] = timestamps
		mu.Unlock()

		return c.Next()
	}
}
